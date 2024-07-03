// Code generated by MockGen. DO NOT EDIT.
// Source: video.go
//
// Generated by this command:
//
//	mockgen -source=video.go -destination=mock/video.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	domain "github.com/raj3k/go-converter/internal/media/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockVideoService is a mock of VideoService interface.
type MockVideoService struct {
	ctrl     *gomock.Controller
	recorder *MockVideoServiceMockRecorder
}

// MockVideoServiceMockRecorder is the mock recorder for MockVideoService.
type MockVideoServiceMockRecorder struct {
	mock *MockVideoService
}

// NewMockVideoService creates a new mock instance.
func NewMockVideoService(ctrl *gomock.Controller) *MockVideoService {
	mock := &MockVideoService{ctrl: ctrl}
	mock.recorder = &MockVideoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVideoService) EXPECT() *MockVideoServiceMockRecorder {
	return m.recorder
}

// DeleteVideo mocks base method.
func (m *MockVideoService) DeleteVideo(ctx context.Context, videoID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVideo", ctx, videoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVideo indicates an expected call of DeleteVideo.
func (mr *MockVideoServiceMockRecorder) DeleteVideo(ctx, videoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVideo", reflect.TypeOf((*MockVideoService)(nil).DeleteVideo), ctx, videoID)
}

// GetVideoByID mocks base method.
func (m *MockVideoService) GetVideoByID(ctx context.Context, videoID string) (*domain.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideoByID", ctx, videoID)
	ret0, _ := ret[0].(*domain.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideoByID indicates an expected call of GetVideoByID.
func (mr *MockVideoServiceMockRecorder) GetVideoByID(ctx, videoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideoByID", reflect.TypeOf((*MockVideoService)(nil).GetVideoByID), ctx, videoID)
}

// UploadVideo mocks base method.
func (m *MockVideoService) UploadVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadVideo", ctx, video, file)
	ret0, _ := ret[0].(*domain.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadVideo indicates an expected call of UploadVideo.
func (mr *MockVideoServiceMockRecorder) UploadVideo(ctx, video, file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadVideo", reflect.TypeOf((*MockVideoService)(nil).UploadVideo), ctx, video, file)
}

// MockVideoRepository is a mock of VideoRepository interface.
type MockVideoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVideoRepositoryMockRecorder
}

// MockVideoRepositoryMockRecorder is the mock recorder for MockVideoRepository.
type MockVideoRepositoryMockRecorder struct {
	mock *MockVideoRepository
}

// NewMockVideoRepository creates a new mock instance.
func NewMockVideoRepository(ctrl *gomock.Controller) *MockVideoRepository {
	mock := &MockVideoRepository{ctrl: ctrl}
	mock.recorder = &MockVideoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVideoRepository) EXPECT() *MockVideoRepositoryMockRecorder {
	return m.recorder
}

// DeleteVideo mocks base method.
func (m *MockVideoRepository) DeleteVideo(ctx context.Context, videoID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVideo", ctx, videoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVideo indicates an expected call of DeleteVideo.
func (mr *MockVideoRepositoryMockRecorder) DeleteVideo(ctx, videoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVideo", reflect.TypeOf((*MockVideoRepository)(nil).DeleteVideo), ctx, videoID)
}

// GetVideoByID mocks base method.
func (m *MockVideoRepository) GetVideoByID(ctx context.Context, videoID string) (*domain.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideoByID", ctx, videoID)
	ret0, _ := ret[0].(*domain.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideoByID indicates an expected call of GetVideoByID.
func (mr *MockVideoRepositoryMockRecorder) GetVideoByID(ctx, videoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideoByID", reflect.TypeOf((*MockVideoRepository)(nil).GetVideoByID), ctx, videoID)
}

// StoreVideo mocks base method.
func (m *MockVideoRepository) StoreVideo(ctx context.Context, video *domain.Video, file io.Reader) (*domain.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreVideo", ctx, video, file)
	ret0, _ := ret[0].(*domain.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreVideo indicates an expected call of StoreVideo.
func (mr *MockVideoRepositoryMockRecorder) StoreVideo(ctx, video, file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreVideo", reflect.TypeOf((*MockVideoRepository)(nil).StoreVideo), ctx, video, file)
}