package redis

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/livegamemodel"
)

type RedisGameInfoUpdatedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      commonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisGameInfoUpdatedSubscriber(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) (commonnotification.NotificationSubscriber[*commonappevent.GameInfoUpdatedAppEvent], error) {
	return &RedisGameInfoUpdatedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisGameInfoUpdatedSubscriber) Subscribe(handler func(*commonappevent.GameInfoUpdatedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewGameInfoUpdatedAppEventChannel(subscriber.liveGameId, subscriber.playerId),
		func(message []byte) {
			event := commonappevent.DeserializeGameInfoUpdatedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
