package rabbitMQ

import (
	"bitbucket.org/kawancicil/core-library/logger"
	"github.com/streadway/amqp"
)

type RabbitMQ interface {
	openConnectionRabbitMQ(dialUrl string) (connection *amqp.Connection, err error)
	closeConnectionRabbitMQ(connection *amqp.Connection)
	failOnError(connection *amqp.Connection, err error, msg string)
	PublishLogger(dialUrl string, payload logger.LoggerPayload)
	ConsumeLogger(dialUrl string, queueName string)(out []logger.Payload)
}
