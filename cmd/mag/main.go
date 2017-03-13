package main

import (
	"log"
	"time"

	common "github.com/ekhabarov/go-common"
	"github.com/mdspinc/mag/agg"
	"github.com/mdspinc/mag/endpoint"
)

func main() {
	cfg := ReadConfig()
	e := endpoint.New(agg.AGGTYPE_STRING, cfg.MaxMessages, time.Duration(cfg.TimeLimit))

	err := e.Listen(cfg.String())
	common.FatalIf(err)
	log.Println("Server quit")
}
