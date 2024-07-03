package domain

import (
	"github.com/google/uuid"
	"time"
)

type Video struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	GridFSFileID string    `json:"grid_fs_file_id"`
	UploadDate   time.Time `json:"upload_date"`
}
