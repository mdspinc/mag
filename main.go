package main

import (
	"log"

	common "github.com/ekhabarov/go-common"
	"github.com/mdspinc/mag/endpoint"
)

func main() {
	cfg := ReadConfig()
	e := endpoint.New()

	err := e.Listen(cfg.String())
	common.FatalIf(err)
	log.Println("Server quit")
}
