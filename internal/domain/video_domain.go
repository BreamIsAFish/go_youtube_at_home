package domain

import (
	databaseModel "go_youtube_at_home/internal/model/database_model"
	requestModel "go_youtube_at_home/internal/model/request"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type VideoController interface {
	PostCreateVideo(ctx *fiber.Ctx) error
	// GetFile(ctx *fiber.Ctx)
	GetVideos(ctx *fiber.Ctx) error
}

type VideoService interface {
	CreateVideo(ctx *fiber.Ctx, req *requestModel.VideoUploadRequest, videoData *multipart.FileHeader, imageData *multipart.FileHeader) error
	// GetVideoByID(ctx *fiber.Ctx) (filePath string, err error)
	// GetFileByID(ctx *fiber.Ctx, id string) (err error)
	// GetVideosByUserID(ctx *fiber.Ctx) ([]*databaseModel.Video, error)
	GetVideos(ctx *fiber.Ctx) ([]*databaseModel.Video, error)
}

type VideoRepository interface {
	CreateVideo(video *databaseModel.Video) error
	// GetVideoByID(videoID string) (filePath string, err error)
	// GetVideosByUserID(userID string) ([]*databaseModel.Video, error)
	GetVideos() ([]*databaseModel.Video, error)
	SaveFileToStorage(filename string, fileData *multipart.FileHeader) (filePath string, err error)
}