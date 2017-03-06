package main

import (
	"errors"
	"strconv"
)

// An Event represents what happened to the telephony system
type Event interface {
}

// Ring means someone is calling and the phone is ringing
type Ring struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Disconnect means a connection has terminated
type Disconnect struct {
	Duration int `json:"duration"`
}

// Connect means a connections was sucessfull (talking)
type Connect struct {
	Extension string `json:"extension"`
	From      string `json:"from"`
}

// Call means an outgoing call is being initiated (not yet connected)
type Call struct {
	Extension string `json:"extension"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func parseEvent(eventType string, fields []string) (Event, error) {
	switch eventType {
	case "RING":
		return parseRing(fields)
	case "DISCONNECT":
		return parseDisconnect(fields)
	case "CONNECT":
		return parseConnect(fields)
	case "CALL":
		return parseCall(fields)
	default:
		return nil, errors.New("Invalid event type")
	}
}

func parseRing(fields []string) (Ring, error) {
	if len(fields) < 2 {
		return Ring{}, errors.New("Not enough fields")
	}
	return Ring{fields[0], fields[1]}, nil
}

func parseDisconnect(fields []string) (Disconnect, error) {
	if len(fields) < 1 {
		return Disconnect{}, errors.New("Not enough fields")
	}

	duration, err := strconv.Atoi(fields[0])
	if err != nil {
		return Disconnect{}, err
	}

	return Disconnect{duration}, nil
}

func parseConnect(fields []string) (Connect, error) {
	if len(fields) < 2 {
		return Connect{}, errors.New("Not enough fields")
	}

	return Connect{fields[0], fields[1]}, nil
}

func parseCall(fields []string) (Call, error) {
	if len(fields) < 3 {
		return Call{}, errors.New("Not enough fields")
	}

	return Call{fields[0], fields[1], fields[2]}, nil
}
