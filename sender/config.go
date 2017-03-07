package sender

import (
	"os"
	"strings"
)

const (
	SLACK_TOKEN   = "SLACK_TOKEN"
	SLACK_CHANNEL = "SLACK_CHANNEL"

	// Include user names into messages with @ prefix
	NOTIFY_USERS = "NOTIFY_USERS"
)

// Configuration settings for Slack sender.
type Config struct {
	token   string
	channel string
	users   []string
}

// Fills out Config struct.
func readConfig() *Config {
	return &Config{
		token:   os.Getenv(SLACK_TOKEN),
		channel: os.Getenv(SLACK_CHANNEL),
		users:   strings.Split(os.Getenv(NOTIFY_USERS), ","),
	}
}
