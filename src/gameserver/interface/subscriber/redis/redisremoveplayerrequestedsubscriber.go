package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type RedisRemovePlayerRequestedSubscriber struct {
	redisProvider *rediseprovider.Provider
}

func NewRedisRemovePlayerRequestedSubscriber() (integrationeventsubscriber.DeprecatedSubscriber[*commonappevent.RemovePlayerRequestedAppEvent], error) {
	return &RedisRemovePlayerRequestedSubscriber{
		redisProvider: rediseprovider.New(),
	}, nil
}

func (subscriber *RedisRemovePlayerRequestedSubscriber) Subscribe(handler func(*commonappevent.RemovePlayerRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewRemovePlayerRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeRemovePlayerRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
