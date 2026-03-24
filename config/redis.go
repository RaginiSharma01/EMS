package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Test-connection
	// 	err := RedisClient.Set(Ctx, "foo", "bar", 0).Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	val, err := RedisClient.Get(Ctx, "foo").Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Redis test value:", val)
}
