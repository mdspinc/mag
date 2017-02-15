package sender

import (
	"errors"

	"github.com/nlopes/slack"
)

var (
	ErrInvalidMessageType = errors.New("sender: slack: invalid message type.")
)

// Slack sender.
type Slack struct {
	client  *slack.Client
	channel string
}

// Initializes Slack sender instance.
func NewSlackSender() (*Slack, error) {
	cfg := readConfig()

	client := slack.New(cfg.token)
	_, err := client.AuthTest()

	if err != nil {
		return nil, err
	}

	return &Slack{
		client:  client,
		channel: cfg.channel,
	}, nil
}

// Sends messages.
func (s *Slack) Send(msg interface{}) error {
	p := slack.NewPostMessageParameters()
	p.AsUser = true

	switch m := msg.(type) {
	case string:
		if _, _, err := s.client.PostMessage(s.channel, m, p); err != nil {
			return err
		}
	default:
		return ErrInvalidMessageType
	}

	return nil
}
