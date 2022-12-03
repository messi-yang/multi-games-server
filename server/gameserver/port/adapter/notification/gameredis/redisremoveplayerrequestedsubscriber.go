package gameredis

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisRemovePlayerRequestedIntegrationEvent struct {
	LiveGameId jsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   jsondto.PlayerIdJsonDto   `json:"playerId"`
}

func NewRedisRemovePlayerRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) RedisRemovePlayerRequestedIntegrationEvent {
	return RedisRemovePlayerRequestedIntegrationEvent{
		LiveGameId: jsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   jsondto.NewPlayerIdJsonDto(playerId),
	}
}

var RedisRemovePlayerRequestedSubscriberChannel string = "remove-player-requested"

type RedisRemovePlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisRemovePlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[RedisRemovePlayerRequestedIntegrationEvent], error) {
	return &RedisRemovePlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisRemovePlayerRequestedSubscriber) Subscribe(handler func(RedisRemovePlayerRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisRemovePlayerRequestedSubscriberChannel, func(message []byte) {
		var event RedisRemovePlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
