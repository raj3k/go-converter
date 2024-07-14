package gridfs

import "go.mongodb.org/mongo-driver/bson"

type GridFS interface {
	UploadFile(filePath, fileName string) error
	DownloadFile(fileName, destPath string) error
	DeleteFile(fileName string) error
	ListFiles() ([]bson.M, error)
	Close() error
}
