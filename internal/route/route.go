package route

import (
	"go_youtube_at_home/internal/controller"
	"go_youtube_at_home/internal/database"
	"go_youtube_at_home/internal/middleware"
	"go_youtube_at_home/internal/repository"
	"go_youtube_at_home/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router) {
	userCollection := database.GetUserCollection()
	userRepository := repository.NewUserRepository(userCollection)
	userService := service.NewUserService(userRepository)
	uc := controller.NewUserController(userService)

	user := api.Group("/user")
	// Setup routes
	user.Post("/", uc.PostNewUser)
	user.Post("/login", uc.PostLogin)
}

func SetupVideoRoutes(api fiber.Router) {
	videoCollection := database.GetVideoCollection()
	videoRepository := repository.NewVideoRepository(videoCollection)
	videoService := service.NewVideoService(videoRepository)
	vc := controller.NewVideoController(videoService)

	video := api.Group("/video")
	// Middleware to check if user is authenticated
	video.Use(middleware.SessionMiddleware)
	// Setup routes
	video.Post("/", vc.PostCreateVideo)
	video.Get("/", vc.GetVideos)
}