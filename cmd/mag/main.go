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
		agg.AggtypeString,
		cfg.MaxMessages,
		time.Duration(cfg.TimeLimit),
	)

	m := monitor.New(
		cfg.APIAddress,
		cfg.APIToken,
		cfg.APITokenRefreshAddress,
		cfg.APITokenRefreshInterval,
		cfg.MonitorInterval,
		cfg.MonitorMaxStoredItems,
		cfg.FKPTreshold,
	)
	m.Start()

	err := e.Listen(cfg.String())
	common.FatalIf(err)
	log.Println("Server quit")
}
