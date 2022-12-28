package redis

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/livegamemodel"
)

type RedisZoomedAreaUpdatedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      commonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisZoomedAreaUpdatedSubscriber(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) (commonnotification.NotificationSubscriber[*commonappevent.ZoomedAreaUpdatedAppEvent], error) {
	return &RedisZoomedAreaUpdatedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisZoomedAreaUpdatedSubscriber) Subscribe(handler func(*commonappevent.ZoomedAreaUpdatedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewZoomedAreaUpdatedAppEventChannel(subscriber.liveGameId, subscriber.playerId),
		func(message []byte) {
			event := commonappevent.DeserializeZoomedAreaUpdatedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
