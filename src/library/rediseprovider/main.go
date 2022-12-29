package rediseprovider

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

type Provider struct {
	redisClient *redis.Client
}

var singleton *Provider

func New() *Provider {
	if singleton == nil {
		singleton = &Provider{
			redisClient: redis.NewClient(&redis.Options{
				Addr:        os.Getenv("REDIS_HOST"),
				Password:    os.Getenv("REDIS_PASSWORD"),
				DB:          0,
				ReadTimeout: -1,
			}),
		}
		return singleton
	}
	return singleton
}

func (subscriber *Provider) Subscribe(channel string, handler func(message []byte)) (unsubscriber func()) {
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

func (subscriber *Provider) Publish(channel string, message []byte) error {
	err := subscriber.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
