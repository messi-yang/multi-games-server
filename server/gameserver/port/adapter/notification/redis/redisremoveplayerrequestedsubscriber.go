package redis

import (
	"encoding/json"

	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
)

type RedisRemovePlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisRemovePlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[commonredisdto.RedisRemovePlayerRequestedEvent], error) {
	return &RedisRemovePlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisRemovePlayerRequestedSubscriber) Subscribe(handler func(commonredisdto.RedisRemovePlayerRequestedEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(commonredisdto.NewRedisRemovePlayerRequestedEventChannel(), func(message []byte) {
		var event commonredisdto.RedisRemovePlayerRequestedEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
