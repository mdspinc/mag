// Package endpoint provides simple TCP listener.
package endpoint

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"

	common "github.com/ekhabarov/go-common"
	"github.com/mdspinc/mag/agg"
	"github.com/mdspinc/mag/handler"
)

type (
	// Handler function for processing messages.
	Handler func(*bufio.ReadWriter, chan interface{})

	// Endpoint represents the TCP listener and connected aggregator.
	Endpoint struct {
		listener net.Listener
		handler  map[string]Handler
		out      chan interface{}
		agg      agg.Aggregator
	}
)

// Initialises new Endpoint instance.
func New(aggType agg.Type) *Endpoint {
	e := &Endpoint{
		handler: map[string]Handler{},
		out:     make(chan interface{}),
	}

	a, err := agg.New(aggType, 5, time.Second*30)
	common.LogIf(err, "endpoint", "new")
	e.agg = a

	e.AddHandler("STR", handler.StringHandler)
	e.AddHandler("ERR", handler.ErrorHandler)
	return e
}

// Add new handler to Endpoint.
func (e *Endpoint) AddHandler(name string, h Handler) {
	e.handler[name] = h
}

// Redirects handled messages to typed aggregator.
func (e *Endpoint) MessageRouter() {
	for {
		v := <-e.out
		switch t := v.(type) {
		case string:
			e.agg.Aggregate(t)
			log.Printf("got: %s; buffer len: %d ", t, e.agg.Count(t))
		case error:
			log.Println("got error value: ", t)
		default:
			log.Println("got unknown type value")
		}
	}
}

// Starts to listen given interface.
func (e *Endpoint) Listen(addr string) (err error) {
	e.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go e.MessageRouter()

	log.Println("Listen on", e.listener.Addr().String())
	for {
		conn, err := e.listener.Accept()
		log.Println("Got connection")
		if err != nil {
			log.Println("Failed accepting a connection request:", err)
			continue
		}
		go e.handleMessages(conn)
	}
}

// Handles incoming messages based on its type.
func (e *Endpoint) handleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	for {
		cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			log.Println("Reached EOF - close this connection.\n   ---")
			return
		case err != nil:
			log.Println("\nError reading command. Got: '"+cmd+"'\n", err)
			return
		}
		cmd = strings.Trim(cmd, "\n ")
		handler, ok := e.handler[cmd]
		if !ok {
			log.Println("Command '" + cmd + "' is not registered.")
			return
		}
		handler(rw, e.out)
	}
}
