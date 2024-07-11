package consumer

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/raj3k/go-converter/pkg/rabbitmq"
	"golang.org/x/sync/errgroup"
)

type consumer struct {
	rabbitmq.ExchangeAndQueueBinding
	workerPoolSize int
	amqpConn       *amqp.Connection
}

var _ EventConsumer = (*consumer)(nil)

func NewConsumer(amqpConn *amqp.Connection, exchangeAndQueueBinding rabbitmq.ExchangeAndQueueBinding, workerPoolSize int) (EventConsumer, error) {
	sub := &consumer{
		ExchangeAndQueueBinding: exchangeAndQueueBinding,
		workerPoolSize:          workerPoolSize,
		amqpConn:                amqpConn,
	}

	return sub, nil
}

func (c *consumer) StartConsumer(ctx context.Context, worker DeliveriesConsumer) error {
	ch, err := c.amqpConn.Channel()
	if err != nil {
		return err
	}

	deliveries, err := ch.Consume(
		c.QueueName,
		c.Consumer,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= c.workerPoolSize; i++ {
		eg.Go(worker(ctx, deliveries, i))
	}

	return eg.Wait()
}
