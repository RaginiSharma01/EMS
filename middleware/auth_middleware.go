package middleware

import (
	"ems/config"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {

	token := c.Get("Authorization")

	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	val, err := config.RedisClient.Get(config.Ctx, token).Result()

	if err == nil && val == "blacklisted" {
		return c.Status(401).JSON(fiber.Map{
			"message": "token invalid, login again",
		})
	}

	return c.Next()
}
