package port

import (
	"context"
	"github.com/raj3k/go-converter/internal/media/core/domain"
	"io"
)

type VideoService interface {
	UploadVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error)
	GetVideoByID(ctx context.Context, videoID string) (*domain.Video, error)
	DeleteVideo(ctx context.Context, videoID string) error
}

type VideoRepository interface {
	StoreVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error)
	GetVideoByID(ctx context.Context, videoID string) (*domain.Video, error)
	DeleteVideo(ctx context.Context, videoID string) error
}
