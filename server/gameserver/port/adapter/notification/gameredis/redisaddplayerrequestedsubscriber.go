package gameredis

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisAddPlayerRequestedIntegrationEvent struct {
	LiveGameId jsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   jsondto.PlayerIdJsonDto   `json:"playerId"`
}

func NewRedisAddPlayerRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) RedisAddPlayerRequestedIntegrationEvent {
	return RedisAddPlayerRequestedIntegrationEvent{
		LiveGameId: jsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   jsondto.NewPlayerIdJsonDto(playerId),
	}
}

var RedisAddPlayerRequestedSubscriberChannel string = "add-player-requested"

type RedisAddPlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisAddPlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[RedisAddPlayerRequestedIntegrationEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAddPlayerRequestedSubscriber) Subscribe(handler func(RedisAddPlayerRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisAddPlayerRequestedSubscriberChannel, func(message []byte) {
		var event RedisAddPlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
