package redis

import (
	"encoding/json"

	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
)

type RedisReviveUnitsRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisReviveUnitsRequestedSubscriber() (commonnotification.NotificationSubscriber[commonredisdto.RedisReviveUnitsRequestedEvent], error) {
	return &RedisReviveUnitsRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisReviveUnitsRequestedSubscriber) Subscribe(handler func(commonredisdto.RedisReviveUnitsRequestedEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(commonredisdto.NewRedisReviveUnitsRequestedEventChannel(), func(message []byte) {
		var event commonredisdto.RedisReviveUnitsRequestedEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
