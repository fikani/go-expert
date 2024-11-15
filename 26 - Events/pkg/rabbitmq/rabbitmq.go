package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

var ConsumerName = "go-consumer"
var QueueName = "go-queue"
var ExchangeName = "amq.direct"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel, nil
}

func Consume(channel *amqp.Channel, out chan<- amqp.Delivery) error {
	messages, err := channel.Consume(
		QueueName,    // queue
		ConsumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return err
	}

	for message := range messages {
		out <- message
	}
	return nil
}

func Publish(channel *amqp.Channel, message string) error {
	err := channel.Publish(
		ExchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	return nil
}
