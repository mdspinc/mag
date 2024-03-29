// Package magclient represents functions for sending messages to server.
package magclient

import (
	"bufio"
	"fmt"
	"net"
)

// Client struct represents client for mag server.
type Client struct {
	conn *net.Conn
	rw   *bufio.ReadWriter
	host string
	port int
}

// NewClient initializes client struct.
func NewClient(host string, port int) (c *Client, err error) {
	c = &Client{}
	c.host = host
	c.port = port
	c.conn, err = c.Connect()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Connect creates new TCP connection to server.
func (c *Client) Connect() (*net.Conn, error) {
	if c.host == "" {
		return nil, ErrHostIsEmpty
	}

	if c.port == 0 {
		return nil, ErrPortIsEmpty
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if err != nil {
		return nil, err
	}

	c.rw, err = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
	if err != nil {
		return nil, err
	}

	return &conn, nil
}

// Reconnect recreate TCP connection.
func (c *Client) Reconnect() (err error) {
	c.rw = nil
	c.conn, err = c.Connect()
	return err
}

// SendString send string data to server.
func (c *Client) SendString(data string) (err error) {
	_, err = c.rw.WriteString("STRING\n")
	if err == nil {
		_, err = c.rw.WriteString(data + "\n")
	}

	if err == nil {
		err = c.rw.Flush()
	}

	return err
}
