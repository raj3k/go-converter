package port

import (
	"context"
	"github.com/google/uuid"
	"github.com/raj3k/go-converter/internal/media/core/domain"
	"io"
)

type VideoService interface {
	UploadVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error)
	GetVideoByID(ctx context.Context, videoID uuid.UUID) (*domain.Video, error)
	DeleteVideo(ctx context.Context, videoID uuid.UUID) error
}

type VideoRepository interface {
	StoreVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error)
	GetVideoByID(ctx context.Context, videoID uuid.UUID) (*domain.Video, error)
	DeleteVideo(ctx context.Context, videoID uuid.UUID) error
}
