package redis

import (
	"encoding/json"

	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
)

type RedisNotificationPublisher struct {
	redisProvider *RedisProvider
}

func NewRedisNotificationPublisher() commonnotification.NotificationPublisher {
	return &RedisNotificationPublisher{
		redisProvider: NewRedisProvider(),
	}
}

func (publisher *RedisNotificationPublisher) Publish(channel string, event any) error {
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = publisher.redisProvider.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
