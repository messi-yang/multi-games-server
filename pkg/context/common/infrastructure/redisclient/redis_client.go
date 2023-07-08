package redisclient

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func Connect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_HOST"),
		DB:          0,
		ReadTimeout: -1,
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}
