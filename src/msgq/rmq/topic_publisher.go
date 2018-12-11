package rmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RMQPublisher struct {
	url               string
	topicExchangeName string
	connection        *amqp.Connection
	channel           *amqp.Channel
	topics            []string
}

func NewRMQPublisher(url string, topicExchangeName string) *RMQPublisher {
	conn, ch := GetChannel(url)
	err := ch.ExchangeDeclare(
		topicExchangeName, // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)

	failOnError(err, "Failed to declare an exchange")
	return &RMQPublisher{
		url:               url,
		topicExchangeName: topicExchangeName,
		connection:        conn,
		channel:           ch,
	}
}

func (pub *RMQPublisher) Publish(msg []byte, destinationTopic string) {
	err := pub.channel.Publish(
		pub.topicExchangeName,
		destinationTopic,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	log.Println("Publish():--topic-exchange-name:", pub.topicExchangeName, ", topic=", destinationTopic)
	log.Printf(" [x] Sent %s", msg)
	log.Println("Publish():err=", err)
	failOnError(err, "Failed to publish a message")
}
