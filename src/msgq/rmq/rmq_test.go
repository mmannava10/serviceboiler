package rmq

import (
	"testing"
	"rmq"
)

func TestRMQ(t *testing.T) {

	url := "amqp://guest:guest@localhost:5672"
	pub := NewRMQPublisher(url, "database.events.trigger_3")

	data1 := "hello"
	data2 := []string{
		"Madhu",
		"Murali",
	}

	msg := EncodeDBMessage(data1, data2)

	pub.publish(msg, "user-data")

}
