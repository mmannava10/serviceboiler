package main

import (
	"common/communications/rmq"
	"constants"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	log.Println("handle(): Got a delivery")
	for d := range deliveries {
		log.Println("handle():Processing a delivery")

		/* 		log.Printf(
			"got %dB delivery: [%v] %s",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		) */
		rmq.DecodeDBMessage(d.Body)

	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}

func main() {

	url := "amqp://guest:guest@localhost:5672"
	sub := rmq.NewRMQSubscriber(url, constants.CONNECT_TOPIC_EXCHANGE)
	log.Println("handle(): Got a delivery")
	sub.Subscribe(constants.CONNECT_DB_TRIGGERS)
	done := make(chan error)
	log.Println("Consume()")
	deliveries := sub.Consume()
	handle(deliveries, done)

	var a string
	fmt.Scanln(&a)
}
