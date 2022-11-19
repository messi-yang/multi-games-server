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

func (service *RedisInfrastructureService) Publish(channel string, message []byte) error {
	err := service.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
