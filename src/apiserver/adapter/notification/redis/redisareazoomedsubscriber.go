package redis

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/livegamemodel"
)

type RedisAreaZoomedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      commonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisAreaZoomedSubscriber(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) (commonnotification.NotificationSubscriber[*commonappevent.AreaZoomedAppEvent], error) {
	return &RedisAreaZoomedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
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
