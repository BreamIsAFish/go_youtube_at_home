package jwt

import (
	"errors"
	"go_youtube_at_home/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type accessClaims[T any] struct {
	Data T `json:"data"`
	jwt.RegisteredClaims
}

func NewAccessClaims[T any](data T) *accessClaims[T] {
	ac := &accessClaims[T]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return ac
}

func (ac *accessClaims[T]) ToToken() (string, error) {
	// claims := customClaims{
	// 	userID,
	// 	jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	},
	// }
	// claims := jwt.MapClaims{
	// 	"userID": userID,                    // Subject (user identifier)
	// 	"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration time
	// 	"iat": time.Now().Unix(),             // Issued at
	// }
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	token, err := jwt.SignedString([]byte(configs.GetConfig().JWT.AccessSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractAccessClaims[T any](token string) (*accessClaims[T], error) {
	jwtToken, err := jwt.ParseWithClaims(token, &accessClaims[T]{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GetConfig().JWT.AccessSecret), nil
	})
	
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*accessClaims[T])
	if !ok {
		return nil, errors.New("Invalid access token claims")
	}
	return claims, nil
}

func (ac *accessClaims[T]) IsExpired() bool {
	now := time.Now()
	exp := ac.ExpiresAt.Time
	return now.After(exp)
}

