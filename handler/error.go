package handler

import (
	"bufio"
	"errors"
	"log"
	"strings"
)

// Hander for Error values.
func ErrorHandler(rw *bufio.ReadWriter, out chan interface{}) {
	log.Print("Receive ERROR message:")
	errMessage, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
	}
	errMessage = strings.Trim(errMessage, "\n ")
	e := errors.New(errMessage)

	//log.Println("Error data:", e)
	out <- e
}
