package gridfs

import (
	"bytes"
	"context"
	"github.com/raj3k/go-converter/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type gridFS struct {
	client *mongo.Client
	db     *mongo.Database
	bucket *gridfs.Bucket
}

type UploadOptions struct {
	Metadata bson.M
}

func (opts *UploadOptions) toGridFSUploadOptions() *options.UploadOptions {
	gridFSUploadOpts := options.GridFSUpload()
	if opts.Metadata != nil {
		gridFSUploadOpts.SetMetadata(opts.Metadata)
	}
	return gridFSUploadOpts
}

func parseUploadOptions(uploadOpts ...*UploadOptions) []*options.UploadOptions {
	var opts []*options.UploadOptions

	for _, uploadOpt := range uploadOpts {
		if uploadOpt == nil {
			continue
		}

		opt := options.UploadOptions{
			Metadata: uploadOpt.Metadata,
		}

		opts = append(opts, &opt)
	}

	return opts
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
	}, nil
}

func (g *gridFS) UploadFile(filePath, fileName string, fileID interface{}, uploadOpts ...*UploadOptions) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var gridFSUploadOpts []*options.UploadOptions

	if uploadOpts != nil {
		gridFSUploadOpts = parseUploadOptions(uploadOpts...)
	}

	uploadStream, err := g.bucket.OpenUploadStreamWithID(fileID, fileName, gridFSUploadOpts...)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (g *gridFS) DownloadFile(destPath string, fileID interface{}) error {
	var buf bytes.Buffer

	_, err := g.bucket.DownloadToStream(fileID, &buf)
	if err != nil {
		return err
	}

	err = os.WriteFile(destPath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g *gridFS) DeleteFile(fileName string) error {
	var fileDoc bson.M
	err := g.db.Collection("fs.files").FindOne(context.Background(), bson.M{"filename": fileName}).Decode(&fileDoc)
	if err != nil {
		return err
	}
	fileID := fileDoc["_id"]

	_, err = g.db.Collection("fs.chunks").DeleteMany(context.Background(), bson.M{"files_id": fileID})
	if err != nil {
		return err
	}

	_, err = g.db.Collection("fs.files").DeleteOne(context.Background(), bson.M{"_id": fileID})
	if err != nil {
		return err
	}

	return nil
}

func (g *gridFS) ListFiles() ([]bson.M, error) {
	cursor, err := g.db.Collection("fs.files").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var files []bson.M
	for cursor.Next(context.Background()) {
		var file bson.M
		if err := cursor.Decode(&file); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (g *gridFS) Close() error {
	return g.client.Disconnect(context.Background())
}
