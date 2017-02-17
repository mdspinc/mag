package handler

import (
	"bufio"
	"errors"
	"log"
	"strings"
)

// Hander for Error values.
func ErrorHandler(rw *bufio.ReadWriter, out chan interface{}) {
	msg, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
		return
	}
	msg = strings.Trim(msg, "\n ")
	out <- errors.New(msg)
}
