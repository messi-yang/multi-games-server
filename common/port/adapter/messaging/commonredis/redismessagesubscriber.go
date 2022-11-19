package commonredis

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/redisclient"
	"github.com/go-redis/redis/v9"
)

type RedisMessageSubscriber struct {
	redisClient *redis.Client
}

func NewRedisMessageSubscriber() *RedisMessageSubscriber {
	return &RedisMessageSubscriber{
		redisClient: redisclient.NewRedisClient(),
	}
}

func (service *RedisMessageSubscriber) Subscribe(channel string, handler func(message []byte)) (unsubscriber func()) {
	pubsub := service.redisClient.Subscribe(context.TODO(), channel)
	go func() {
		for msg := range pubsub.Channel() {
			handler([]byte(msg.Payload))
		}
	}()

	return func() {
		pubsub.Close()
	}
}
