package redisclient

import (
	"os"

	"github.com/go-redis/redis/v9"
)

var redisClient *redis.Client

func NewRedisClient() *redis.Client {
	if redisClient == nil {

		redisClient = redis.NewClient(&redis.Options{
			Addr:        os.Getenv("REDIS_HOST"),
			Password:    os.Getenv("REDIS_PASSWORD"),
			DB:          0,
			ReadTimeout: -1,
		})
		return redisClient
	}
	return redisClient
}
