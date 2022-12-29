package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type RedisZoomAreaRequestedSubscriber struct {
	redisProvider *rediseprovider.Provider
}

func NewRedisZoomAreaRequestedSubscriber() (integrationeventsubscriber.DeprecatedSubscriber[*commonappevent.ZoomAreaRequestedAppEvent], error) {
	return &RedisZoomAreaRequestedSubscriber{
		redisProvider: rediseprovider.New(),
	}, nil
}

func (subscriber *RedisZoomAreaRequestedSubscriber) Subscribe(handler func(*commonappevent.ZoomAreaRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewZoomAreaRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeZoomAreaRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
