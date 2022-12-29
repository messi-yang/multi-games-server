package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type RedisAreaZoomedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      commonmodel.PlayerId
	redisProvider *rediseprovider.Provider
}

func NewRedisAreaZoomedSubscriber(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) (integrationeventsubscriber.DeprecatedSubscriber[*commonappevent.AreaZoomedAppEvent], error) {
	return &RedisAreaZoomedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: rediseprovider.New(),
	}, nil
}

func (subscriber *RedisAreaZoomedSubscriber) Subscribe(handler func(*commonappevent.AreaZoomedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewAreaZoomedAppEventChannel(subscriber.liveGameId, subscriber.playerId),
		func(message []byte) {
			event := commonappevent.DeserializeAreaZoomedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
