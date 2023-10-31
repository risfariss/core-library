package rabbitMQ

import (
	"bitbucket.org/kawancicil/core-library/logger"
	"encoding/json"
	"github.com/streadway/amqp"
)

func (r *RabbitMQUtils) PublishLogger(dialUrl string, payload logger.LoggerPayload){
	connection, err := r.openConnectionRabbitMQ(dialUrl)
	if err != nil {
		return
	}

	channel, err := connection.Channel()
	r.failOnError(connection, err, "failed to open a channel")

	q, err := channel.QueueDeclare(
		payload.Name, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	r.failOnError(connection, err, "failed to declare a queue")

	body, _ := json.Marshal(payload.Payload)
	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	r.failOnError(connection, err, "failed to publish a message")
	r.closeConnectionRabbitMQ(connection)
	return
}
