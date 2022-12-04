package redis

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
)

type RedisAreaZoomedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      gamecommonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisAreaZoomedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonnotification.NotificationSubscriber[commonredisdto.RedisAreaZoomedEvent], error) {
	return &RedisAreaZoomedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAreaZoomedSubscriber) Subscribe(handler func(commonredisdto.RedisAreaZoomedEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(commonredisdto.NewRedisAreaZoomedEventChannel(subscriber.liveGameId, subscriber.playerId), func(message []byte) {
		var event commonredisdto.RedisAreaZoomedEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
