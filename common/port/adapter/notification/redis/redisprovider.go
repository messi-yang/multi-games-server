package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type RedisProvider struct {
	redisClient *redis.Client
}

func NewRedisProvider() *RedisProvider {
	return &RedisProvider{
		redisClient: NewRedisClient(),
	}
}

func (subscriber *RedisProvider) Subscribe(channel string, handler func(message []byte)) (unsubscriber func()) {
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

func (subscriber *RedisProvider) Publish(channel string, message []byte) error {
	err := subscriber.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
