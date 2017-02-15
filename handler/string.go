package handler

import (
	"bufio"
	"log"
	"strings"
)

// Handler for String values
func StringHandler(rw *bufio.ReadWriter, out chan interface{}) {
	log.Print("Receive STRING message:")
	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
	}
	out <- strings.Trim(s, "\n ")

	_, err = rw.WriteString("Response message.\n")
	if err != nil {
		log.Println("string handler: response write error:", err)
	}

	if err = rw.Flush(); err != nil {
		log.Println("string handler: flush error:", err)
	}
}
