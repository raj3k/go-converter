package publisher

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type publisher struct {
	log      *log.Logger
	amqpConn *amqp.Connection
	amqpChan *amqp.Channel
}

var _ EventPublisher = (*publisher)(nil)

func NewPublisher(amqpConn *amqp.Connection, log *log.Logger) (*publisher, error) {
	channel, err := amqpConn.Channel()
	if err != nil {
		return nil, err
	}

	return &publisher{
		log:      log,
		amqpChan: channel,
		amqpConn: amqpConn,
	}, nil
}

func (p *publisher) Publish(ctx context.Context, exchangeName, bidingKey string, msg amqp.Publishing) error {
	ch, err := p.amqpConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	log.Info("Publishing message to exchange", "exchange", exchangeName, "key", bidingKey)

	if err = ch.PublishWithContext(
		ctx,
		exchangeName,
		bidingKey,
		true,
		true,
		msg,
	); err != nil {
		return err
	}

	return nil
}
