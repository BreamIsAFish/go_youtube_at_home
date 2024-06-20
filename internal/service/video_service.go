package service

import (
	"go_youtube_at_home/internal/domain"
	databaseModel "go_youtube_at_home/internal/model/database_model"
	requestModel "go_youtube_at_home/internal/model/request"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type videoService struct {
	videoRepository domain.VideoRepository
}

func NewVideoService(videoRepository domain.VideoRepository) domain.VideoService {
	return &videoService{
		videoRepository,
	}
}

func (s *videoService) CreateVideo(ctx *fiber.Ctx, req *requestModel.VideoUploadRequest, videoData *multipart.FileHeader, imageData *multipart.FileHeader) error {
	// Get the userID from the context
	userID := ctx.Locals("userID").(string);
	// Save video to storage
	videoID := uuid.New().String()
	imageID := uuid.New().String()
	videoPath, err := s.videoRepository.SaveFileToStorage(videoID, videoData)
	if err != nil {
		return err
	}
	imagePath, err := s.videoRepository.SaveFileToStorage(imageID, imageData)
	if err != nil {
		return err
	}

	video := &databaseModel.Video{
		ID: videoID,
		Description: req.Description,
		Title: req.Title,
		URL: videoPath,
		ThumbnailURL: imagePath,
		UploaderID: userID,
	}
	if err = s.videoRepository.CreateVideo(video); err != nil {
		return err
	}
	
	return nil
}

// func (s *videoService) GetVideosByUserID(ctx *fiber.Ctx) ([]*databaseModel.Video, error) {
// 	userID := ctx.Locals("userID").(string)
// 	videos, err := s.videoRepository.GetVideosByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return videos, nil
// }

func (s *videoService) GetVideos(ctx *fiber.Ctx) ([]*databaseModel.Video, error) {
	videos, err := s.videoRepository.GetVideos()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

