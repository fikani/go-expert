package main

import (
	"events/pkg/rabbitmq"
	"os"
)

func main() {
	// make sure argument is passed
	if len(os.Args) < 2 {
		panic("message is required")
	}

	message := os.Args[1]
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	err = rabbitmq.Publish(channel, message)
	if err != nil {
		panic(err)
	}
}
