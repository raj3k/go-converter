package domain

import "time"

type Video struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	GridFSFileID string    `json:"grid_fs_file_id"`
	UploadDate   time.Time `json:"upload_date"`
}
