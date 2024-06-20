package domain

import (
	databaseModel "go_youtube_at_home/internal/model/database_model"
	requestModel "go_youtube_at_home/internal/model/request"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	PostNewUser(ctx *fiber.Ctx) error
	PostLogin(ctx *fiber.Ctx) error
}

type UserService interface {
	CreateUser(req requestModel.UserRegisterRequest) (string, error)
	Login(req requestModel.UserLoginRequest) (string, error)
}

type UserRepository interface {
	CreateUser(user *databaseModel.User) error
	GetUserByUsername(username string) (*databaseModel.User, error)
}
