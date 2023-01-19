package redispub

import (
	"context"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/go-redis/redis/v9"
)

type redisPublisher struct {
	redisClient *redis.Client
}

func New() intgrevent.IntgrEventPublisher {
	return &redisPublisher{
		redisClient: redis.NewClient(&redis.Options{
			Addr:        os.Getenv("REDIS_HOST"),
			Password:    os.Getenv("REDIS_PASSWORD"),
			DB:          0,
			ReadTimeout: -1,
		}),
	}
}

func (publisher *redisPublisher) Publish(channel string, message []byte) error {
	err := publisher.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
