package middleware

import (
	"ems/config"
	"ems/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "authorization header missing",
		})
	}

	// Extract token
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Verify JWT
	_, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	// Check Redis blacklist
	val, err := config.RedisClient.Get(config.Ctx, token).Result()
	if err == nil && val == "blacklisted" {
		return c.Status(401).JSON(fiber.Map{
			"message": "token invalid, login again",
		})
	}

	return c.Next()
}
