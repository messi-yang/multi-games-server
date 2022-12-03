package apiredis

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisGameInfoUpdatedIntegrationEvent struct {
	LiveGameId jsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   jsondto.PlayerIdJsonDto   `json:"playerId"`
	Dimension  jsondto.DimensionJsonDto  `json:"dimension"`
}

func NewRedisGameInfoUpdatedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, dimension gamecommonmodel.Dimension) RedisGameInfoUpdatedIntegrationEvent {
	return RedisGameInfoUpdatedIntegrationEvent{
		LiveGameId: jsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   jsondto.NewPlayerIdJsonDto(playerId),
		Dimension:  jsondto.NewDimensionJsonDto(dimension),
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

func NewRedisGameInfoUpdatedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonnotification.NotificationSubscriber[RedisGameInfoUpdatedIntegrationEvent], error) {
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
