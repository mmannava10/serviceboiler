package rmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RMQSubscriber struct {
	url          string
	exchangeName string
	connection   *amqp.Connection
	channel      *amqp.Channel
	queue        *amqp.Queue
	topics       []string
}

// func comments
//

func NewRMQSubscriber(url string, exchangeName string) *RMQSubscriber {
	conn, ch := GetChannel(url)
	q := GetQueue(exchangeName, ch)
	log.Println("NewRMQSubscriber():exchange-name=", exchangeName, " queue-name=", q.Name)
	return &RMQSubscriber{
		exchangeName: exchangeName,
		url:          url,
		connection:   conn,
		channel:      ch,
		queue:        q,
	}
}

func (sub *RMQSubscriber) Subscribe(toTopic string) {
	log.Println("Subscribe():exchange-name=", sub.exchangeName, ", queue-name=", sub.queue.Name, ", topic=", toTopic)

	// TODO: if subscribed alreday skip
	err := sub.channel.QueueBind(
		sub.queue.Name,   // queue name
		toTopic,          // routing key
		sub.exchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind to a queue")
}

func (sub *RMQSubscriber) Close() {
	err := sub.channel.Close()
	failOnError(err, "Failed close channel")

	err = sub.connection.Close()
	failOnError(err, "Failed close connection to exchange")
}

func (sub *RMQSubscriber) Consume() <-chan amqp.Delivery {

	log.Println("Consume():queue-name=", sub.queue.Name)

	deliveries, err := sub.channel.Consume(
		sub.queue.Name, // queue
		"",             // consumer
		false,          // auto ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")
	return deliveries
}
