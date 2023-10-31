package rabbitMQ

import (
	"bitbucket.org/kawancicil/core-library/logger"
	"encoding/json"
	"log"
)

func (r *RabbitMQUtils) ConsumeLogger(dialUrl string, queueName string)(out []logger.Payload){
	connection, err := r.openConnectionRabbitMQ(dialUrl)
	if err != nil {
		return
	}

	channel, err := connection.Channel()
	r.failOnError(connection, err, "failed to open a channel")

	data, err := channel.Consume(
		queueName, // name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	r.failOnError(connection, err, "failed to declare a consumer")

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	for message := range data {
		var dataMessage logger.Payload
		// For example, show received message in a console.
		log.Printf(" > Received message: %s\n", message.Body)
		err = json.Unmarshal(message.Body,&dataMessage)
		out = append(out,dataMessage)
	}

	return
}

