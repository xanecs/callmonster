package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// A Message represents emitted by the call monitor.
type Message struct {
	Timestamp    time.Time `json:"timestamp"`
	ConnectionID int       `json:"connectionId"`
	Type         string    `json:"type"`
	Event        Event     `json:"event"`
}

// A Parser parses string messages
type Parser struct {
	location *time.Location
}

// NewParser creates and returns a new Parser
func NewParser(config *tomlConfig) (Parser, error) {
	location, err := time.LoadLocation(config.Callmonitor.Timezone)
	if err != nil {
		return Parser{}, err
	}
	return Parser{location}, nil
}

// Parse takes a string line from the call monitor and returns a Message object
func (p Parser) Parse(input string) (Message, error) {
	fields := strings.Split(input, ";")
	fields = fields[:len(fields)-1]

	if len(fields) < 3 {
		return Message{}, errors.New("Invalid number of fields in received string")
	}
	dateStr := fields[0]
	eventType := fields[1]
	connectionStr := fields[2]

	// TODO: Replace UTC with configurable timezone
	timestamp, err := time.ParseInLocation("02.01.06 15:04:05", dateStr, p.location)
	if err != nil {
		return Message{}, err
	}

	connectionID, err := strconv.Atoi(connectionStr)
	if err != nil {
		return Message{}, err
	}

	event, err := parseEvent(eventType, fields[3:])
	if err != nil {
		return Message{}, err
	}

	return Message{
		Timestamp:    timestamp,
		ConnectionID: connectionID,
		Type:         eventType,
		Event:        event,
	}, nil

}
