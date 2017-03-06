package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type tomlConfig struct {
	Callmonitor struct {
		Host     string
		Timezone string
	}
	Redis struct {
		Host  string
		Topic string
	}
}

func loadConfig(filename string) (tomlConfig, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return tomlConfig{}, errors.New("Specified file does not exist")
	}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return tomlConfig{}, nil
	}
	var config tomlConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		return tomlConfig{}, err
	}
	return config, nil
}
