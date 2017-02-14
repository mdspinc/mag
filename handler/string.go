package handler

import (
	"bufio"
	"log"
	"strings"

	common "github.com/ekhabarov/go-common"
)

func StringHandler(rw *bufio.ReadWriter, out chan interface{}) {
	log.Print("Receive STRING message:")
	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
	}
	s = strings.Trim(s, "\n ")
	out <- s
	_, err = rw.WriteString("Response message.\n")
	if err != nil {
		common.FatalIf(err)
	}
	err = rw.Flush()
	if err != nil {
		common.FatalIf(err)
	}
}
