package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient *redis.Client

var SMTPEmail string
var SMTPPassword string

func ConnectRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func LoadSMTPConfig() {
	SMTPEmail = os.Getenv("SMTP_EMAIL")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
}
