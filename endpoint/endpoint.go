package endpoint

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type (
	Handler func(*bufio.ReadWriter)

	Endpoint struct {
		listener net.Listener
		handler  map[string]Handler
	}
)

func New() *Endpoint {
	return &Endpoint{handler: map[string]Handler{}}
}

func (e *Endpoint) AddHandler(name string, h Handler) {
	e.handler[name] = h
}

func (e *Endpoint) Listen(addr string) error {
	var err error
	e.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("Listen on", e.listener.Addr().String())
	for {
		log.Println("Accept a connection request.")
		conn, err := e.listener.Accept()
		if err != nil {
			log.Println("Failed accepting a connection request:", err)
			continue
		}
		//log.Println("Handle incoming messages.")
		go e.handleMessage(conn)
	}
}

func (e *Endpoint) handleMessage(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	for {
		//log.Print("Receive command '")
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
		//log.Println(cmd + "'")
		handler, ok := e.handler[cmd]
		if !ok {
			log.Println("Command '" + cmd + "' is not registered.")
			return
		}
		handler(rw)
	}
}
