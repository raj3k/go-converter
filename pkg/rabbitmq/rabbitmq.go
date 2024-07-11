package rabbitmq

import (
	"context"
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

type ExchangeAndQueueBinding struct {
	ExchangeName string `mapstructure:"exchangeName" validate:"required"`
	ExchangeKind string `mapstructure:"exchangeKind" validate:"required"`
	QueueName    string `mapstructure:"queueName" validate:"required"`
	BindingKey   string `mapstructure:"bindingKey" validate:"required"`
	Concurrency  int    `mapstructure:"concurrency" validate:"required"`
	Consumer     string `mapstructure:"consumer" validate:"required"`
}

type Config struct {
	URI string `mapstructure:"uri" validate:"required"`
}

func NewRabbitMQConn(cfg Config) (*amqp.Connection, error) {
	var (
		amqpConn *amqp.Connection
		counts   int64
	)

	for {
		connection, err := amqp.Dial(cfg.URI)
		if err != nil {
			log.Error("failed to connect to RabbitMQ...", err, cfg.URI)
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

func DeclareBinding(ctx context.Context, channel *amqp.Channel, exchangeAndQueueBinding ExchangeAndQueueBinding) (amqp.Queue, error) {
	if err := DeclareExchange(ctx, channel, exchangeAndQueueBinding.ExchangeName, exchangeAndQueueBinding.ExchangeKind); err != nil {
		return amqp.Queue{}, err
	}

	queue, err := DeclareQueue(ctx, channel, exchangeAndQueueBinding.QueueName)
	if err != nil {
		return amqp.Queue{}, err
	}

	if err := BindQueue(ctx, channel, queue.Name, exchangeAndQueueBinding.BindingKey, exchangeAndQueueBinding.ExchangeName); err != nil {
		return amqp.Queue{}, err
	}

	return queue, nil
}

func DeclareQueue(ctx context.Context, channel *amqp.Channel, name string) (amqp.Queue, error) {
	return channel.QueueDeclare(name, true, false, false, false, nil)
}

func DeclareExchange(ctx context.Context, channel *amqp.Channel, name, kind string) error {
	return channel.ExchangeDeclare(name, kind, true, false, false, false, nil)
}

func BindQueue(ctx context.Context, channel *amqp.Channel, queue, key, exchange string) error {
	return channel.QueueBind(queue, key, exchange, false, nil)
}
