package main

import (
	"log"

	common "github.com/ekhabarov/go-common"
	"github.com/mdspinc/mag/agg"
	"github.com/mdspinc/mag/endpoint"
)

func main() {
	cfg := ReadConfig()
	e := endpoint.New(agg.AGGTYPE_STRING) //, sender.NewSlackSender())

	err := e.Listen(cfg.String())
	common.FatalIf(err)
	log.Println("Server quit")
}
