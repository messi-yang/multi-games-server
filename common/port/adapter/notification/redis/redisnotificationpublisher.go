package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/notification"
)

type RedisNotificationPublisher struct {
	redisProvider *RedisProvider
}

func NewRedisNotificationPublisher() notification.NotificationPublisher {
	return &RedisNotificationPublisher{
		redisProvider: NewRedisProvider(),
	}
}

func (publisher *RedisNotificationPublisher) Publish(channel string, jsonMessage any) error {
	message, err := json.Marshal(jsonMessage)
	if err != nil {
		return err
	}

	err = publisher.redisProvider.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
