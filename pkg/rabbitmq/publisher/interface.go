package publisher

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventPublisher interface {
	Publish(ctx context.Context, exchange, key string, msg amqp.Publishing) error
}
