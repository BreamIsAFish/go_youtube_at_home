package middleware

import (
	"go_youtube_at_home/internal/model"
	"go_youtube_at_home/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
    // Get the Authorization header
    token := c.Get("Authorization")
    if token == "" {
        // Handle missing Authorization header
				c.SendStatus(fiber.StatusUnauthorized)
        return nil
    }
		claims, err := jwt.ExtractAccessClaims[model.VideoSession](token)
		if err != nil {
			c.SendStatus(fiber.StatusUnauthorized)
			return nil
		}
		c.Locals("userID", claims.Data.UserID)
		c.Next()
		return nil
}