package main

import (
	"fmt"
	"os"

	common "github.com/ekhabarov/go-common"
)

const (
	LISTEN_ADDRESS   = "LISTEN_ADDRESS"
	LISTEN_PORT      = "LISTEN_PORT"
	AGG_MAX_MESSAGES = "AGG_MAX_MESSAGES"
	AGG_TIME_LIMIT   = "AGG_TIME_LIMIT"
)

// Configuration settings.
type Config struct {
	Address     string
	Port        int
	MaxMessages int
	TimeLimit   int
}

// Returns address with port to listen to.
func (c *Config) String() string {
	return fmt.Sprintf("%s:%d", c.Address, c.Port)
}

// Fills out Config struct.
func ReadConfig() *Config {
	cfg := &Config{
		Address: os.Getenv(LISTEN_ADDRESS),
	}
	common.ReadEnvIntParam(&cfg.Port, 3050, LISTEN_PORT)
	common.ReadEnvIntParam(&cfg.MaxMessages, 30, AGG_MAX_MESSAGES)
	common.ReadEnvIntParam(&cfg.TimeLimit, 30, AGG_TIME_LIMIT)

	return cfg
}
