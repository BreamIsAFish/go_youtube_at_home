package controller

import (
	"go_youtube_at_home/internal/domain"
	requestModel "go_youtube_at_home/internal/model/request"

	Validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type videoController struct {
	videoService domain.VideoService
	validator *Validator.Validate
}

func NewVideoController(videoService domain.VideoService) domain.VideoController {
	return &videoController{
		videoService: videoService,
		validator: Validator.New(),
	}
}

func (vc *videoController) PostCreateVideo(ctx *fiber.Ctx) error {
	var req requestModel.VideoUploadRequest
	ctx.BodyParser(&req)

	// Parse title and description from the form
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error parsing request",
		})
		return err
	}

	// Validate request
	if err := vc.validator.Struct(req); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	video, err := ctx.FormFile("video_data")
	if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Error retrieving video",
			})
			return err
	}
	image, err := ctx.FormFile("thumbnail_data")
	if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Error retrieving thumbnail image",
			})
			return err
	}

	
	if err := vc.videoService.CreateVideo(ctx, &req, video, image); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	ctx.Status(201).JSON(fiber.Map{
		"message": "Video uploaded",
	})
	// ctx.Status(201).JSON(fiber.Map{
	// 	"message": "Passed",
	// 	//req
	// 	"data":	fiber.Map{
	// 		"video": req,
	// 	},
	// })
	return nil
}

func (vc *videoController) GetVideos(ctx *fiber.Ctx) error {
	videos, err := vc.videoService.GetVideos(ctx)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	ctx.Status(200).JSON(fiber.Map{
		"message": "Successfully logged in",
		"data": fiber.Map{
			"videos": videos,
		},
	})
	return nil
}

// func (vc *videoController) GetVideosByUserID(ctx *fiber.Ctx) {
// 	videos, err := vc.videoService.GetVideosByUserID(ctx)
// 	if err != nil {
// 		ctx.Status(400).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	// ctx.Status(200).JSON(fiber.Map{
// 	// 	"message": "Videos fetched",
// 	// 	"data": fiber.Map{
// 	// 		"videos": videos,
// 	// 	},
// 	// })
// 	ctx.Status(200).JSON(fiber.Map{
// 		"message": videos,
// 	})
// }
