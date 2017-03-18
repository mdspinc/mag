// Package endpoint provides simple TCP listener.
package endpoint

import (
	"bufio"
	"fmt"
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

// New initialises new Endpoint instance.
func New(aggType agg.Type, maxMessages int, timeLimit time.Duration) *Endpoint {
	e := &Endpoint{
		handler: map[string]Handler{},
		out:     make(chan interface{}),
	}

	a, err := agg.New(aggType, maxMessages, time.Second*timeLimit)
	common.LogIf(err, "endpoint: new")
	e.agg = a

	e.AddHandler("STRING", handler.StringHandler)
	e.AddHandler("ERROR", handler.ErrorHandler)
	return e
}

// AddHandler adds new handler to Endpoint.
func (e *Endpoint) AddHandler(name string, h Handler) {
	e.handler[name] = h
}

// MessageRouter redirects handled messages to typed aggregator.
func (e *Endpoint) MessageRouter() {
	for {
		v := <-e.out
		switch t := v.(type) {
		case string:
			e.agg.Aggregate(t)
		case error:
			//log.Println("got error value: ", t)
		default:
			log.Println("endpoint: message router: unknown type value")
		}
	}
}

// Listen starts to listen given interface.
func (e *Endpoint) Listen(addr string) (err error) {
	e.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go e.MessageRouter()

	log.Println("Listen on", e.listener.Addr().String())
	for {
		conn, err := e.listener.Accept()
		if err != nil {
			log.Println("endpoint: listen:", err)
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
			//log.Println("Client distconnected.")
			return
		case err != nil:
			//log.Println("\nError reading command. Got: '"+cmd+"'\n", err)
			return
		}
		cmd = strings.Trim(cmd, "\n ")
		handler, ok := e.handler[cmd]
		if !ok {
			s := fmt.Sprintf("Command %q is not registered.\n", cmd)
			log.Println(s)
			send(rw, s)
			return
		}
		handler(rw, e.out)
	}
}
