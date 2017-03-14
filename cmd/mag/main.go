package main

import (
	"log"
	"time"

	common "github.com/ekhabarov/go-common"
	"github.com/mdspinc/mag/agg"
	"github.com/mdspinc/mag/endpoint"
	"github.com/mdspinc/mag/monitor"
)

func main() {
	cfg := ReadConfig()
	e := endpoint.New(
		agg.AGGTYPE_STRING,
		cfg.MaxMessages,
		time.Duration(cfg.TimeLimit),
	)

	m := monitor.New(
		cfg.BotsmetricsApiAddress,
		cfg.MonitorInterval,
		cfg.MonitorMaxStoredItems,
	)
	m.Start()

	err := e.Listen(cfg.String())
	common.FatalIf(err)
	log.Println("Server quit")
}
