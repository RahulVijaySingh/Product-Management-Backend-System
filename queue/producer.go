package queue

import (
	"log"
	"github.com/streadway/amqp"
)

func PublishToQueue(queueName, message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.Publish("", queueName, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})
	if err != nil {
		return err
	}

	log.Println("Message published to queue:", queueName)
	return nil
}