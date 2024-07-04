package rabbitmq

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	retryTimes     = 5
	backOffSeconds = 2
)

var ErrCannotConnectRabbitMQ = errors.New("cannot connect to RabbitMQ")

func NewRabbitMQConn(rabbitMQURL string) (*amqp.Connection, error) {
	var (
		amqpConn *amqp.Connection
		counts   int64
	)

	for {
		connection, err := amqp.Dial(rabbitMQURL)
		if err != nil {
			log.Error("failed to connect to RabbitMQ...", err, rabbitMQURL)
			counts++
		} else {
			amqpConn = connection
			break
		}

		if counts > retryTimes {
			log.Error("failed to retry", err)
			return nil, ErrCannotConnectRabbitMQ
		}

		log.Info("Backing off for 2 seconds...")
		time.Sleep(backOffSeconds * time.Second)
	}

	log.Info("connected to rabbitmq 🎉")

	return amqpConn, nil
}
