package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/mdspinc/mag/endpoint"
)

func StringHandler(rw *bufio.ReadWriter) {
	log.Print("Receive STRING message:")
	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
	}
	s = strings.Trim(s, "\n ")
	log.Println(s)
	_, err = rw.WriteString("Response message.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}

func main() {
	e := endpoint.New()
	e.AddHandler("STR", StringHandler)

	err := e.Listen()
	FatalIf(err)
	log.Println("Server quit")
}
