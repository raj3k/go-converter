package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/raj3k/go-converter/internal/media/core/domain"
	"github.com/raj3k/go-converter/internal/media/core/port"
	"io"
)

type VideoService struct {
	videoRepo port.VideoRepository
}

func NewVideoService(videoRepo port.VideoRepository) *VideoService {
	return &VideoService{
		videoRepo: videoRepo,
	}
}

func (vs *VideoService) UploadVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error) {
	return nil, nil
}

func (vs *VideoService) GetVideoByID(ctx context.Context, video uuid.UUID) (*domain.Video, error) {
	return nil, nil
}

func (vs *VideoService) DeleteVideo(ctx context.Context, videoID uuid.UUID) error {
	return nil
}
