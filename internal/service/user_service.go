package service

import (
	"crypto/rand"
	"errors"
	"go_youtube_at_home/internal/domain"
	"go_youtube_at_home/internal/model"
	databaseModel "go_youtube_at_home/internal/model/database_model"
	requestModel "go_youtube_at_home/internal/model/request"
	"go_youtube_at_home/pkg/jwt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository,
	}
}

func (s *userService) CreateUser(req requestModel.UserRegisterRequest) (string, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return "", err
	}
	user := &databaseModel.User{
		ID: uuid.New().String(),
		Username: req.Username,
		Password: hashedPassword,
	}
	err = s.userRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (s *userService) Login(req requestModel.UserLoginRequest) (string, error) {
	user, err := s.userRepository.GetUserByUsername(req.Username)
	if err != nil {
		return "", err
	}
	if !comparePasswords(user.Password, req.Password) {
		return "", errors.New("Invalid password")
	}
	claims := jwt.NewAccessClaims(&model.VideoSession{
		UserID: user.ID,
	})
	token, err := claims.ToToken()
	if err != nil {
		return "", err
	}
	return token, nil
}

func comparePasswords(hashedPassword string, inputPassword string) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(inputPassword))
	return err == nil
}

func generateSalt() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	return string(bytes), err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}



