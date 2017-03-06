package main

import (
	"bufio"
	"net"
)

// Callmon represents a connection to a callmon server
type Callmon struct {
	host string
}

// NewCallmon creates and returns a new Callmon object
func NewCallmon(config *tomlConfig) Callmon {
	return Callmon{config.Callmonitor.Host}
}

// Connect connects to the call monitor server and returns a channel for messages
func (c *Callmon) Connect(chMsg chan string, chErr chan error) error {
	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		return err
	}
	defer conn.Close()

	msgReader := bufio.NewReader(conn)

	for {
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			chErr <- err
			continue
		}
		chMsg <- msg
	}
}
