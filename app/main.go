package main

import (
	"fmt"
	"go_youtube_at_home/configs"
	"go_youtube_at_home/internal/database"
	"go_youtube_at_home/internal/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	if err := configs.SetupViperConfig(); err != nil {
		panic(err)
	}
	if err := database.InitMongo(); err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()	
	app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
			AllowHeaders: "*",
		}),
	)

	app.Static("/", "./storage")

	api := app.Group("/api")

	// Setup routes
	route.SetupUserRoutes(api)
	route.SetupVideoRoutes(api)
	
	if err := app.Listen(fmt.Sprintf(":%d", configs.GetConfig().App.Port)); err != nil {
		panic(err)
	}
} 