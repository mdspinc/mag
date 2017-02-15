package sender

type (
	Type int

	Transport interface {
		Send(interface{}) error
	}
)
