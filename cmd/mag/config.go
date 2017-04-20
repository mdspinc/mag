package main

import (
	"fmt"
	"os"

	common "github.com/ekhabarov/go-common"
)

const (
	// Listen for incoming messages on address
	listenAddress = "LISTEN_ADDRESS"

	// Port for incoming messages
	listenPort = "LISTEN_PORT"

	// Number of messages to store. If limit excceeded, send all stored messaeges.
	aggMaxMessages = "AGG_MAX_MESSAGES"

	// If limit of messages is not exceeded, send existing messages if its
	// more than zero every AGG_TIME_LIMIT seconds.
	aggTimeLimit = "AGG_TIME_LIMIT"

	// API address
	APIAddress = "API_ADDRESS"

	// API JWT token
	APIToken = "API_TOKEN"

	// Number of seconds betwenn requests for token refresh.
	APITokenRefreshInterval = "API_TOKEN_REFRESH_INTERVAL"

	// Address for refresh token request.
	APITokenRefreshAddress = "API_TOKEN_REFRESH_ADDRESS"

	// Number of seconds between requests BOTSMETRIC_API_ADDRESS
	monitorInterval = "MONITOR_INTERVAL"

	// Maximum number of requests to store for analyze.
	monitorMaxStoredItems = "MONITOR_MAX_STORED_ITEMS"

	// Send notification if number of errors of type FREQUENCY_KEY_PRESENT
	// for last MONITOR_INTERVAL seconds is less than threshold value.
	fkpThreshold = "FKP_THRESHOLD"
)

// Config is a configuration settings store.
type Config struct {
	Address                 string
	Port                    int
	MaxMessages             int
	TimeLimit               int
	APIAddress              string
	APIToken                string
	APITokenRefreshInterval int
	APITokenRefreshAddress  string
	MonitorInterval         int
	MonitorMaxStoredItems   int
	FKPTreshold             int
}

// Returns address with port to listen to.
func (c *Config) String() string {
	return fmt.Sprintf("%s:%d", c.Address, c.Port)
}

// ReadConfig fills out Config struct.
func ReadConfig() *Config {
	cfg := &Config{
		Address:                os.Getenv(listenAddress),
		APIAddress:             os.Getenv(APIAddress),
		APIToken:               os.Getenv(APIToken),
		APITokenRefreshAddress: os.Getenv(APITokenRefreshAddress),
	}

	common.ReadEnvIntParam(&cfg.Port, 3050, listenPort)
	common.ReadEnvIntParam(&cfg.MaxMessages, 30, aggMaxMessages)
	common.ReadEnvIntParam(&cfg.TimeLimit, 30, aggTimeLimit)
	common.ReadEnvIntParam(&cfg.MonitorInterval, 300, monitorInterval)
	common.ReadEnvIntParam(&cfg.MonitorMaxStoredItems, 3, monitorMaxStoredItems)
	common.ReadEnvIntParam(&cfg.FKPTreshold, 1000, fkpThreshold)
	common.ReadEnvIntParam(&cfg.APITokenRefreshInterval, 86400, APITokenRefreshInterval)

	return cfg
}
