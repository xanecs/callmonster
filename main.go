package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := loadConfig("config.toml")
	if err != nil {
		panic(err)
	}
	callmon := NewCallmon(&config)
	parser, err := NewParser(&config)
	if err != nil {
		panic(err)
	}
	broker := NewBroker(&config)

	chMsg := make(chan string)
	chErr := make(chan error)
	go callmon.Connect(chMsg, chErr)
	for {
		select {
		case msg := <-chMsg:
			message, err := parser.Parse(msg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				continue
			}
			if err = broker.Send(message); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case err := <-chErr:
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
