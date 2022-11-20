package commonredisnotification

import (
	"context"
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/notification"
	"github.com/go-redis/redis/v9"
)

type RedisNotificationPublisher struct {
	redisClient *redis.Client
}

func NewRedisNotificationPublisher() notification.NotificationPublisher {
	return &RedisNotificationPublisher{
		redisClient: redisclient.NewRedisClient(),
	}
}

func (publisher *RedisNotificationPublisher) Publish(channel string, jsonMessage any) error {
	message, err := json.Marshal(jsonMessage)
	if err != nil {
		return err
	}

	err = publisher.redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
