package infrastructureservice

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

type RedisInfrastructureService interface {
	Subscribe(channel string, handler func(message []byte)) (unsubscriber func())
	Publish(channel string, message []byte) error
}

type redisInfrastructureService struct {
	redisClient *redis.Client
}

var redisClient *redis.Client

func NewRedisInfrastructureService() RedisInfrastructureService {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:        os.Getenv("REDIS_HOST"),
			Password:    os.Getenv("REDIS_PASSWORD"),
			DB:          0,
			ReadTimeout: -1,
		})
	}
	return &redisInfrastructureService{
		redisClient: redisClient,
	}
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
