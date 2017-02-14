package main

import (
	"bufio"
	"log"
	"strings"

	common "github.com/ekhabarov/go-common"
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
		common.FatalIf(err)
	}
	err = rw.Flush()
	if err != nil {
		common.FatalIf(err)
	}
}

func main() {
	cfg := ReadConfig()
	e := endpoint.New()
	e.AddHandler("STR", StringHandler)

	err := e.Listen(cfg.String())
	common.FatalIf(err)
	log.Println("Server quit")
}
