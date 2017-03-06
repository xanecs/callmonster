package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		assert.FailNow(t, "Can't get cwd")
	}
	filename, err := filepath.Abs(filepath.Join(cwd, "fixtures/config1.toml"))
	if err != nil {
		assert.FailNow(t, "Can't get absolute path of config fixture")
	}
	config, err := loadConfig(filename)
	assert.Nil(t, err, "Error occured: %v", err)

	assert.Equal(t, "10.0.0.1:1012", config.Callmonitor.Host, "Invalid call monitor host")
	assert.Equal(t, "UTC", config.Callmonitor.Timezone, "Invalid timezone")
	assert.Equal(t, "10.0.0.223", config.Redis.Host, "Invalid redis host")
	assert.Equal(t, "call", config.Redis.Topic, "Invalid redis topic")
}

func TestLoadConfigErrors(t *testing.T) {
	var err error
	_, err = loadConfig("fixtures/doesnotexist.toml")
	assert.Error(t, err, "Accepted invalid filename")

	cwd, err := os.Getwd()
	if err != nil {
		assert.FailNow(t, "Can't get cwd")
	}
	filename, err := filepath.Abs(filepath.Join(cwd, "fixtures/configInvalid.toml"))
	if err != nil {
		assert.FailNow(t, "Can't get absolute path of config fixture")
	}

	_, err = loadConfig(filename)
	assert.Error(t, err, "Accepted invalid toml")
}
