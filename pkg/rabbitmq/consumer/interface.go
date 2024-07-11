package consumer

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type DeliveriesConsumer func(ctx context.Context, deliveries <-chan amqp.Delivery, workerID int) func() error

type EventConsumer interface {
	StartConsumer(ctx context.Context, worker DeliveriesConsumer) error
}
