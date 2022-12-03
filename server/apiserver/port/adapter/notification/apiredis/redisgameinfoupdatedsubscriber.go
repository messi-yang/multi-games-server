package apiredis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/module/common/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/presenter/presenterdto"
)

type RedisGameInfoUpdatedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
	Dimension  presenterdto.DimensionPresenterDto  `json:"dimension"`
}

func NewRedisGameInfoUpdatedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, dimension gamecommonmodel.Dimension) RedisGameInfoUpdatedIntegrationEvent {
	return RedisGameInfoUpdatedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
		Dimension:  presenterdto.NewDimensionPresenterDto(dimension),
	}
}

func RedisGameInfoUpdatedSubscriberChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisGameInfoUpdatedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      gamecommonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisGameInfoUpdatedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (notification.NotificationSubscriber[RedisGameInfoUpdatedIntegrationEvent], error) {
	return &RedisGameInfoUpdatedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisGameInfoUpdatedSubscriber) Subscribe(handler func(RedisGameInfoUpdatedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisGameInfoUpdatedSubscriberChannel(subscriber.liveGameId, subscriber.playerId), func(message []byte) {
		var event RedisGameInfoUpdatedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
