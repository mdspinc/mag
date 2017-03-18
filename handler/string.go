package handler

import (
	"bufio"
	"log"
	"strings"
)

// StringHandler is a handler for String values.
func StringHandler(rw *bufio.ReadWriter, out chan interface{}) {
	msg, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
		return
	}
	out <- strings.Trim(msg, "\n ")
}
