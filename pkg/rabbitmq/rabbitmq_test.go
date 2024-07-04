package rabbitmq_test

import (
	"github.com/raj3k/go-converter/pkg/rabbitmq"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNewRabbitMQConn(t *testing.T) {
	rabbiMQURL := "amqp://guest:guest@localhost:5672"

	conn, err := rabbitmq.NewRabbitMQConn(rabbiMQURL)
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	log.Printf("Successfully connected to RabbitMQ and opened a channel")
}
