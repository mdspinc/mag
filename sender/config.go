package sender

import "os"

const (
	SLACK_TOKEN   = "SLACK_TOKEN"
	SLACK_CHANNEL = "SLACK_CHANNEL"
)

type Config struct {
	token   string
	channel string
}

func readConfig() *Config {
	return &Config{
		token:   os.Getenv(SLACK_TOKEN),
		channel: os.Getenv(SLACK_CHANNEL),
	}
}
