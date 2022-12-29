package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type RedisBuildItemRequestedSubscriber struct {
	redisProvider *rediseprovider.Provider
}

func NewRedisBuildItemRequestedSubscriber() (integrationeventsubscriber.DeprecatedSubscriber[*commonappevent.BuildItemRequestedAppEvent], error) {
	return &RedisBuildItemRequestedSubscriber{
		redisProvider: rediseprovider.New(),
	}, nil
}

func (subscriber *RedisBuildItemRequestedSubscriber) Subscribe(handler func(*commonappevent.BuildItemRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewBuildItemRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeBuildItemRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
