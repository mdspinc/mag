package main

import (
	"fmt"
	"os"

	common "github.com/ekhabarov/go-common"
)

const (
	LISTEN_ADDRESS = "LISTEN_ADDRESS"
	LISTEN_PORT    = "LISTEN_PORT"
)

// Configuration settings.
type Config struct {
	Address string
	Port    int
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

	return cfg
}
