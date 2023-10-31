package rabbitMQ

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQUtils struct {
}

func InitRabbitMQUtils() RabbitMQ {
	return &RabbitMQUtils{
	}
}

func (r *RabbitMQUtils) openConnectionRabbitMQ(dialUrl string) (connection *amqp.Connection, err error) {
	fmt.Println("try start RabbitMQ ...")
	connection, err = amqp.Dial(dialUrl)
	if err != nil {
		//r.failOnError(connection, err, "failed to connect to RabbitMQ")
		err = errors.New("failed to connect to RabbitMQ")
		log.Println("error openConnectionRabbitMQ Dial Url :",dialUrl)
		return
	}

	return connection, nil
}

func (r *RabbitMQUtils) closeConnectionRabbitMQ(connection *amqp.Connection) {
	fmt.Println("try close RabbitMQ ...")
	_ = connection.Close()
	return
}

func (r *RabbitMQUtils) failOnError(connection *amqp.Connection, err error, msg string) {
	if err != nil {
		_ = connection.Close()
		log.Println(msg, err.Error())
	}
	return
}
