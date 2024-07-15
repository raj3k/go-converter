package gridfs

import "go.mongodb.org/mongo-driver/bson"

type GridFS interface {
	UploadFile(filePath, fileName string, fileID interface{}, uploadOpts ...*UploadOptions) error
	DownloadFile(destPath string, fileID interface{}) error
	DeleteFile(fileName string) error
	ListFiles() ([]bson.M, error)
	Close() error
}
