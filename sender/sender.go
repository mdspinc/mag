package sender

type (
	// Transport for messages.
	Transport interface {
		Send(interface{}) error
	}
)
