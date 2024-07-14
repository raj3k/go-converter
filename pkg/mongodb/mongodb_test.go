package mongodb_test

import (
	"context"
	"github.com/raj3k/go-converter/pkg/mongodb"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewMongoDBConn(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg := mongodb.Config{
		URI:      "mongodb://localhost:27017",
		User:     "root",
		Password: "example",
	}

	mongoDBConn, err := mongodb.NewMongoDBConn(ctx, cfg)
	if err != nil {
		t.Error(err)
	}
	defer mongoDBConn.Disconnect(ctx)

	log.Println("Connected to MongoDB")
}
