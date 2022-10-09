package infrastructureservice

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

type RedisInfrastructureService interface {
	Subscribe(channel string, handler func(message []byte)) (unsubscriber func())
	Publish(channel string, message []byte) error
	Set(key string, value []byte) error
	Get(key string) (value []byte, err error)
}

type redisInfrastructureService struct {
	redisClient *redis.Client
}

var redisInfrastructureServiceInstance *redisInfrastructureService

func NewRedisInfrastructureService() RedisInfrastructureService {
	if redisInfrastructureServiceInstance == nil {
		return &redisInfrastructureService{
			redisClient: redis.NewClient(&redis.Options{
				Addr:        os.Getenv("REDIS_HOST"),
				Password:    os.Getenv("REDIS_PASSWORD"),
				DB:          0,
				ReadTimeout: -1,
			}),
		}
	}
	return redisInfrastructureServiceInstance
}

func (service *redisInfrastructureService) Subscribe(channel string, handler func(message []byte)) (unsubscriber func()) {
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

func (service *redisInfrastructureService) Publish(channel string, message []byte) error {
	err := service.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (service *redisInfrastructureService) Set(key string, value []byte) error {
	err := service.redisClient.Set(context.TODO(), key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (service *redisInfrastructureService) Get(key string) (value []byte, err error) {
	val, err := service.redisClient.Get(context.TODO(), key).Result()
	if err != nil {
		return []byte{}, err
	}

	return []byte(val), nil
}
