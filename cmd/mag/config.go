package main

import (
	"fmt"
	"os"

	common "github.com/ekhabarov/go-common"
)

const (
	// Listen for incoming messages on address
	LISTEN_ADDRESS = "LISTEN_ADDRESS"

	// Port for incoming messages
	LISTEN_PORT = "LISTEN_PORT"

	// Number of messages to store. If limit excceeded, send all stored messaeges.
	AGG_MAX_MESSAGES = "AGG_MAX_MESSAGES"

	// If limit of messages is not exceeded, send existing messages if its
	// more than zero every AGG_TIME_LIMIT seconds.
	AGG_TIME_LIMIT = "AGG_TIME_LIMIT"

	// Botsmetrics API address
	BOTSMETRICS_API_ADDRESS = "BOTSMETRICS_API_ADDRESS"

	// Number of seconds between requests BOTSMETRIC_API_ADDRESS
	MONITOR_INTERVAL = "MONITOR_INTERVAL"

	// Maximum number of requests to store for analyze.
	MONITOR_MAX_STORED_ITEMS = "MONITOR_MAX_STORED_ITEMS"
)

// Configuration settings.
type Config struct {
	Address               string
	Port                  int
	MaxMessages           int
	TimeLimit             int
	BotsmetricsApiAddress string
	MonitorInterval       int
	MonitorMaxStoredItems int
}

// Returns address with port to listen to.
func (c *Config) String() string {
	return fmt.Sprintf("%s:%d", c.Address, c.Port)
}

// Fills out Config struct.
func ReadConfig() *Config {
	cfg := &Config{
		Address:               os.Getenv(LISTEN_ADDRESS),
		BotsmetricsApiAddress: os.Getenv(BOTSMETRICS_API_ADDRESS),
	}

	common.ReadEnvIntParam(&cfg.Port, 3050, LISTEN_PORT)
	common.ReadEnvIntParam(&cfg.MaxMessages, 30, AGG_MAX_MESSAGES)
	common.ReadEnvIntParam(&cfg.TimeLimit, 30, AGG_TIME_LIMIT)
	common.ReadEnvIntParam(&cfg.MonitorInterval, 300, MONITOR_INTERVAL)
	common.ReadEnvIntParam(&cfg.MonitorMaxStoredItems, 3, MONITOR_MAX_STORED_ITEMS)

	return cfg
}
