package sender

import (
	"errors"
	"strings"

	"github.com/nlopes/slack"
)

var (
	// ErrInvalidMessageType happen when invalidd message type is presented.
	ErrInvalidMessageType = errors.New("sender: slack: invalid message type")
)

// Slack sender.
type Slack struct {
	client  *slack.Client
	channel string
	// List of users with ID for @mention with default messages
	mention string

	// List of user for monitoring
	monitorMention string
}

// MessageType is a type of sended messages, could be "default" or "monitor".
type MessageType int

const (
	// DefaultMessage is a message type by default.
	DefaultMessage MessageType = iota
	// MonitorMessage is a message type for monitoring notifications.
	MonitorMessage
)

// NewSlackSender initializes Slack sender instance.
func NewSlackSender() (*Slack, error) {
	cfg := readConfig()

	client := slack.New(cfg.token)
	_, err := client.AuthTest()
	if err != nil {
		return nil, err
	}

	m := []string{}
	mm := []string{}
	slackUsers, err := client.GetUsers()
	if err != nil {
		return nil, err
	}

	for _, u := range cfg.users {
		for _, su := range slackUsers {
			if su.Name == u {
				m = append(m, u)
			}
		}
	}

	for _, u := range cfg.monitorUsers {
		for _, su := range slackUsers {
			if su.Name == u {
				mm = append(mm, u)
			}
		}
	}

	return &Slack{
		client:         client,
		channel:        cfg.channel,
		mention:        "@" + strings.Join(m, ", @"),
		monitorMention: "@" + strings.Join(mm, ", @"),
	}, nil
}

// Send sends messages.
func (s *Slack) Send(msg interface{}, msgType MessageType) error {
	p := slack.NewPostMessageParameters()
	p.AsUser = true
	p.LinkNames = 1

	switch m := msg.(type) {
	case string:
		switch msgType {
		case MonitorMessage:
			m += s.monitorMention
		default:
			m += s.mention
		}

		if _, _, err := s.client.PostMessage(s.channel, m, p); err != nil {
			return err
		}
	default:
		return ErrInvalidMessageType
	}

	return nil
}
