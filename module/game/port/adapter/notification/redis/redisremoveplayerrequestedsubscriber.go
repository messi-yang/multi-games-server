package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/module/common/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/presenter/presenterdto"
)

type RedisRemovePlayerRequestedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
}

func NewRedisRemovePlayerRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) RedisRemovePlayerRequestedIntegrationEvent {
	return RedisRemovePlayerRequestedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
	}
}

var RedisRemovePlayerRequestedSubscriberChannel string = "remove-player-requested"

type RedisRemovePlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisRemovePlayerRequestedSubscriber() (notification.NotificationSubscriber[RedisRemovePlayerRequestedIntegrationEvent], error) {
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
