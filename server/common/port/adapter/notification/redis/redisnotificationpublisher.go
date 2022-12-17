package redis

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type RedisNotificationPublisher struct {
	redisProvider *RedisProvider
}

func NewRedisNotificationPublisher() commonnotification.NotificationPublisher {
	return &RedisNotificationPublisher{
		redisProvider: NewRedisProvider(),
	}
}

func (publisher *RedisNotificationPublisher) Publish(channel string, event event.AppEvent) error {
	message := event.Serialize()

	err := publisher.redisProvider.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
