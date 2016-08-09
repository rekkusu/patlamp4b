package patlamp

import (
	"bytes"
	"errors"
	"net"
)

type Client struct {
	Connection *net.TCPConn
}

// LED/Sound Command Pattern
const (
	LED_RED          = 1
	LED_YELLOW       = 2
	LED_GREEN        = 4
	BUZZER_SHORT     = 8
	BUZZER_LONG      = 16
	LED_RED_BLINK    = 32
	LED_YELLOW_BLINK = 64
	LED_GREEN_BLINK  = 128
)

func (c Client) Connect(host string) error {
	ip, err := net.ResolveTCPAddr("ip4", host)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("ip4", nil, ip)
	if err != nil {
		return err
	}
	c.Connection = conn
	return nil
}

func (c Client) WriteState(mode byte) error {
	if c.Connection == nil {
		return errors.New("Connection is not established")
	}

	cmd := []byte{'W', mode}
	c.Connection.Write(cmd)

	resp := make([]byte, 2)
	_, err := c.Connection.Read(resp)
	if err != nil {
		return err
	}
	if bytes.Equal([]byte("ACK"), resp) {
		return nil
	}
	return nil
}

func (c Client) ReadState() (byte, error) {
	if c.Connection == nil {
		return 0, errors.New("Connection is not established")
	}

	cmd := []byte{'R'}
	c.Connection.Write(cmd)

	resp := make([]byte, 2)
	_, err := c.Connection.Read(resp)
	if err != nil {
		return 0, err
	}
	return resp[1], nil
}
