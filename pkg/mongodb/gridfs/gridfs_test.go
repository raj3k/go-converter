package gridfs_test

import (
	"context"
	"github.com/raj3k/go-converter/pkg/mongodb"
	"github.com/raj3k/go-converter/pkg/mongodb/gridfs"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewGridFSBucket(t *testing.T) {
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

	bucket, err := gridfs.NewGridFSBucket(mongoDBConn.Database("files"))
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		t.Error(err)
	}

	uploadStream, err := bucket.OpenUploadStream("example.txt")
	if err != nil {
		t.Error(err)
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data)
	if err != nil {
		t.Error(err)
	}

	log.Println("File uploaded successfully")
}
