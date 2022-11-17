package service

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/redisclient"
	"github.com/go-redis/redis/v9"
)

type RedisInfrastructureService struct {
	redisClient *redis.Client
}

var RedisInfrastructureServiceInstance *RedisInfrastructureService

func NewRedisInfrastructureService() *RedisInfrastructureService {
	if RedisInfrastructureServiceInstance == nil {
		return &RedisInfrastructureService{
			redisClient: redisclient.NewRedisClient(),
		}
	}
	return RedisInfrastructureServiceInstance
}

func (service *RedisInfrastructureService) Subscribe(channel string, handler func(message []byte)) (unsubscriber func()) {
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

func (service *RedisInfrastructureService) Publish(channel string, message []byte) error {
	err := service.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (service *RedisInfrastructureService) Set(key string, value []byte) error {
	err := service.redisClient.Set(context.TODO(), key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (service *RedisInfrastructureService) Get(key string) (value []byte, err error) {
	val, err := service.redisClient.Get(context.TODO(), key).Result()
	if err != nil {
		return []byte{}, err
	}

	return []byte(val), nil
}
