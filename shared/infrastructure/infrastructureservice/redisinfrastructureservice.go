package infrastructureservice

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

type RedisService interface {
	Subscribe(channel string, handler func(message []byte)) (unsubscriber func())
	Publish(channel string, message []byte) error
	Set(key string, value []byte) error
	Get(key string) (value []byte, err error)
}

type redisService struct {
	redisClient *redis.Client
}

var redisServiceInstance *redisService

func NewRedisService() RedisService {
	if redisServiceInstance == nil {
		return &redisService{
			redisClient: redis.NewClient(&redis.Options{
				Addr:        os.Getenv("REDIS_HOST"),
				Password:    os.Getenv("REDIS_PASSWORD"),
				DB:          0,
				ReadTimeout: -1,
			}),
		}
	}
	return redisServiceInstance
}

func (service *redisService) Subscribe(channel string, handler func(message []byte)) (unsubscriber func()) {
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

func (service *redisService) Publish(channel string, message []byte) error {
	err := service.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (service *redisService) Set(key string, value []byte) error {
	err := service.redisClient.Set(context.TODO(), key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (service *redisService) Get(key string) (value []byte, err error) {
	val, err := service.redisClient.Get(context.TODO(), key).Result()
	if err != nil {
		return []byte{}, err
	}

	return []byte(val), nil
}
