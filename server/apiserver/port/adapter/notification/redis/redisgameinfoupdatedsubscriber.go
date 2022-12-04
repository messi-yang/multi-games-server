package redis

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisGameInfoUpdatedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      gamecommonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisGameInfoUpdatedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonnotification.NotificationSubscriber[*commonapplicationevent.GameInfoUpdatedApplicationEvent], error) {
	return &RedisGameInfoUpdatedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisGameInfoUpdatedSubscriber) Subscribe(handler func(*commonapplicationevent.GameInfoUpdatedApplicationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonapplicationevent.NewGameInfoUpdatedApplicationEventChannel(subscriber.liveGameId, subscriber.playerId),
		func(message []byte) {
			event := commonapplicationevent.DeserializeGameInfoUpdatedApplicationEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
