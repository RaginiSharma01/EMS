package middleware

import (
	"fmt"
	"time"

	"ems/config"

	"github.com/gofiber/fiber/v3"
)

func LoginRateLimiter(c fiber.Ctx) error {

	ip := c.IP()

	key := fmt.Sprintf("login_rate:%s", ip)

	// increment request count
	count, err := config.RedisClient.Incr(config.Ctx, key).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Redis error",
		})
	}

	// set expiration on first request
	if count == 1 {
		config.RedisClient.Expire(config.Ctx, key, time.Minute)
	}

	// allow only 5 requests per minute
	if count > 5 {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Too many login attempts. Try again later.",
		})
	}

	return c.Next()
}
