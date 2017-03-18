package sender

import (
	"os"
	"strings"
)

const (
	slackToken   = "SLACK_TOKEN"
	slackChannel = "SLACK_CHANNEL"

	// Include user names into messages with @ prefix
	notifyUsers        = "NOTIFY_USERS"
	notifyMonitorUsers = "NOTIFY_MONITOR_USERS"
)

// Config contains settings for Slack sender.
type Config struct {
	token        string
	channel      string
	users        []string
	monitorUsers []string
}

// readConfig fills out Config struct.
func readConfig() *Config {
	return &Config{
		token:        os.Getenv(slackToken),
		channel:      os.Getenv(slackChannel),
		users:        strings.Split(os.Getenv(notifyUsers), ","),
		monitorUsers: strings.Split(os.Getenv(notifyMonitorUsers), ","),
	}
}
