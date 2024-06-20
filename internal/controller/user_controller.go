package controller

import (
	"go_youtube_at_home/internal/domain"
	requestModel "go_youtube_at_home/internal/model/request"

	Validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService domain.UserService
	validator		*Validator.Validate
}

func NewUserController(userService domain.UserService) domain.UserController {
	return &userController{
		userService: userService,
		validator: Validator.New(),
	}
}

func (uc *userController) PostNewUser(ctx *fiber.Ctx) error {
	var req requestModel.UserRegisterRequest
	ctx.BodyParser(&req)

	err := uc.validator.Struct(req)
	if err != nil {
		// for _, err := range err.(Validator.ValidationErrors) {
		// 		var el IError
		// 		el.Field = err.Field()
		// 		el.Tag = err.Tag()
		// 		el.Value = err.Param()
		// 		errors = append(errors, &el)
		// }
		// c.SendStatus(fiber.StatusBadRequest).JSON(errors)
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}
	// req := &requestModel.UserRegisterRequest{
	// 	Username: ctx.Params("username"),
	// 	Password: ctx.Params("password"),
	// }
	userID, err := uc.userService.CreateUser(req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	ctx.Status(201).JSON(fiber.Map{
		"message": "User created",
		"data": fiber.Map{
			"user_id": userID,
		}})
	return nil
}

func (uc *userController) PostLogin(ctx *fiber.Ctx) error  {
	var req requestModel.UserLoginRequest
	ctx.BodyParser(&req)

	// Validate request
	err := uc.validator.Struct(req)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	token, err := uc.userService.Login(req)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}
	ctx.Status(200).JSON(fiber.Map{
		"message": "Successfully logged in",
		"data": fiber.Map{
			"token": token,
		},
	})
	return nil
}