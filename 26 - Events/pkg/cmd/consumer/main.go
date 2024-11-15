package main

import (
	"events/pkg/rabbitmq"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	incomming := make(chan amqp.Delivery)

	go rabbitmq.Consume(channel, incomming)

	for message := range incomming {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
