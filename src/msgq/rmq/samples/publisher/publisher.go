package main

import (
	"constants"
	"fmt"
	"msgq/rmq"
)

func main() {

	url := constants.RMQ_URL
	pub := rmq.NewRMQPublisher(url, constants.CONNECT_TOPIC_EXCHANGE)

	data1 := "user-address"
	data2 := []string{
		"1511",
		"Daphne",
	}
	msg := rmq.EncodeDBMessage(data1, data2)
	pub.Publish(msg, constants.CONNECT_DB_TRIGGERS)
	var a string
	fmt.Scanln(&a)
}
