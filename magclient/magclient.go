package magclient

import (
	"errors"
	"fmt"
	"net"
)

var (
	client *Client

	// ErrHostIsEmpty happen when host is empty string or it's not presented.
	ErrHostIsEmpty = errors.New("magclient: setup: host is empty")

	// ErrPortIsEmpty happen when port is not presented.
	ErrPortIsEmpty = errors.New("magclient: setup: port is empty")

	// ErrClientIsNotInitialized happen if client was not be initialized.
	ErrClientIsNotInitialized = errors.New("magclient: setup: client is not initialized")
)

// Setup initializes client variable.
func Setup(host string, port int) (err error) {
	client, err = NewClient(host, port)
	return err
}

// Send sends data to server and handle reconnects.
func Send(data string) error {
	if client == nil || client.rw == nil {
		return ErrClientIsNotInitialized
	}

	err := client.SendString(data)
	if err != nil {
		if nerr, ok := err.(net.Error); ok {
			if rerr := client.Reconnect(); rerr != nil {
				return fmt.Errorf("magclient: send: reconnect error: %s\n", rerr)
			}
			return fmt.Errorf("magclient: send: network error: %s\n", nerr)
		}
		return fmt.Errorf("magclient: send: error: %s\n", err)
	}

	return nil
}
