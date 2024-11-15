package producer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const url = "amqp://guest:guest@rabbitmq:5672/"

type ProducerST struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func DialProducer() (*ProducerST, error) {

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"command-control", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &ProducerST{Connection: conn, Channel: ch, Queue: q}, nil
}
