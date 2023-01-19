package redissub

import (
	"context"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/go-redis/redis/v9"
)

type redisSubscriber struct {
	redisClient *redis.Client
}

func New() intevent.IntEventSubscriber {
	return &redisSubscriber{
		redisClient: redis.NewClient(&redis.Options{
			Addr:        os.Getenv("REDIS_HOST"),
			Password:    os.Getenv("REDIS_PASSWORD"),
			DB:          0,
			ReadTimeout: -1,
		}),
	}
}

func (subscriber *redisSubscriber) Subscribe(channel string, handler func([]byte)) func() {
	pubsub := subscriber.redisClient.Subscribe(context.TODO(), channel)
	go func() {
		for msg := range pubsub.Channel() {
			handler([]byte(msg.Payload))
		}
	}()

	return func() {
		pubsub.Close()
	}
}
