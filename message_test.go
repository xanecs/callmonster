package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	config, err := loadConfig("fixtures/config1.toml")
	if err != nil {
		assert.FailNow(t, "Could not open config")
	}
	p, err := NewParser(&config)
	if err != nil {
		assert.FailNow(t, "Could not create parser")
	}

	msg1, err := p.Parse("05.03.17 18:16:28;RING;0;0162834568;83456;SIP2;")
	assert.Nil(t, err, "Error occured: %v", err)
	assert.Equal(t, 0, msg1.ConnectionID, "Wrong connection id")
	assert.True(t, msg1.Timestamp.Equal(time.Date(2017, 03, 05, 18, 16, 28, 0, time.UTC)), "Invalid timestamp, got: %v", msg1.Timestamp)
	ringEvent, ok := msg1.Event.(Ring)
	assert.True(t, ok, "Wrong event type")
	assert.Equal(t, "0162834568", ringEvent.From, "Wrong From number")
	assert.Equal(t, "83456", ringEvent.To, "Wrong To number")

	msg2, err := p.Parse("05.03.17 18:16:32;DISCONNECT;51;0;")
	assert.Nil(t, err, "Error occured: %v", err)
	assert.Equal(t, 51, msg2.ConnectionID, "Wrong connection id")
	assert.True(t, msg2.Timestamp.Equal(time.Date(2017, 03, 05, 18, 16, 32, 0, time.UTC)), "Invalid timestamp, got: %v", msg2.Timestamp)
	disconnectEvent, ok := msg2.Event.(Disconnect)
	assert.True(t, ok, "Wrong event type")
	assert.Equal(t, 0, disconnectEvent.Duration, "Wrong duration")
}

func TestParseErrors(t *testing.T) {
	var err error

	config, err := loadConfig("fixtures/config1.toml")
	if err != nil {
		assert.FailNow(t, "Could not open config")
	}
	p, err := NewParser(&config)
	if err != nil {
		assert.FailNow(t, "Could not create parser")
	}

	_, err = p.Parse("05.03.17 18:16:28;RING;")
	assert.Error(t, err, "Accepted message with not enough fields")

	_, err = p.Parse("05.03.17 18:16:28;RING;0;0162834568;")
	assert.Error(t, err, "Accepted message with not enough fields")

	_, err = p.Parse("05.03.17 18:16:28;CONNECT;0;0162834568;")
	assert.Error(t, err, "Accepted message with not enough fields")

	_, err = p.Parse("05.03.17 18:16:28;DISCONNECT;0;")
	assert.Error(t, err, "Accepted message with not enough fields")

	_, err = p.Parse("05.03.17 18:16:28;DISCONNECT;0;notanint;")
	assert.Error(t, err, "Accepted disconnect message with invalid duration")

	_, err = p.Parse("05.03.17 18:16:28;CALL;0;0123;")
	assert.Error(t, err, "Accepted message with not enough fields")

	_, err = p.Parse("05.03.17 18:16:28;OTHER;0;")
	assert.Error(t, err, "Accepted invalid field type")

	_, err = p.Parse("05.03.17 18:16:28;DISCONNECT;not_an_int;0;")
	assert.Error(t, err, "Accepted message with invalid connection id")

	_, err = p.Parse("30.02.17 18:16:28;DISCONNECT;0;0;")
	assert.Error(t, err, "Accepted message with invalid date")
}
