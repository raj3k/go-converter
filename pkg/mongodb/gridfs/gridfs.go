package gridfs

import (
	"context"
	"github.com/raj3k/go-converter/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type gridFS struct {
	client *mongo.Client
	db     *mongo.Database
	bucket *gridfs.Bucket
	ctx    context.Context
}

var _ GridFS = (*gridFS)(nil)

func NewGridFSBucket(ctx context.Context, config mongodb.Config) (GridFS, error) {
	mongoConn, err := mongodb.NewMongoDBConn(ctx, config)
	if err != nil {
		return nil, err
	}

	db := mongoConn.Database(config.Database)
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return nil, err
	}

	return &gridFS{
		client: mongoConn,
		db:     db,
		bucket: bucket,
		ctx:    ctx,
	}, nil
}

func (g gridFS) UploadFile(filePath, fileName string) error {
	//TODO implement me
	panic("implement me")
}

func (g gridFS) DownloadFile(fileName, destPath string) error {
	//TODO implement me
	panic("implement me")
}

func (g gridFS) DeleteFile(fileName string) error {
	//TODO implement me
	panic("implement me")
}

func (g gridFS) ListFiles() ([]bson.M, error) {
	//TODO implement me
	panic("implement me")
}

func (g gridFS) Close() error {
	//TODO implement me
	panic("implement me")
}
