// Package sender represents connectors for different services such as Slack, EMail, etc.
package sender

type (
	// Transport for messages.
	Transport interface {
		Send(interface{}, MessageType) error
	}
)
