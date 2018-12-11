package rmq

import (
	"fmt"
	"log"
	"pb/db"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establoish connection to message broker")

	ch, err := conn.Channel()
	failOnError(err, "Failed to get a channel for the connection")
	return conn, ch
}

func InitExchange(topic_exchange string, ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		topic_exchange, // name
		"topic",        // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)

	failOnError(err, "Failed to declare an exchange")
}

func GetQueue(topic_exchange string, ch *amqp.Channel) *amqp.Queue {
	InitExchange(topic_exchange, ch)

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare queue on the channel")
	return &q
}

func DecodeDBMessage(theEncodedMessage []byte) *database.DatabaseTrigger {
	dbTrigger := &database.DatabaseTrigger{}
	if err := proto.Unmarshal(theEncodedMessage, dbTrigger); err != nil {
		log.Fatalln("Failed to parse Database Trigger:", err)
	}

	log.Printf("Received a message: Trigger=", dbTrigger.Trigger, ", Data1=", dbTrigger.Data1, ", Data2=", dbTrigger.Data2)
	return dbTrigger
}

func EncodeDBMessage(data1 string, data2 []string) []byte {
	log.Println("EncodeDBMessage(data1=", data1, ", data2=", data2)
	dbTrigger := &pb_db.DatabaseTrigger{
		Trigger: database.DBTrigger_TRIGGER_1,
		Data1:   data1,
		Data2:   data2,
	}

	msg, err := proto.Marshal(dbTrigger)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	return msg
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}

}
