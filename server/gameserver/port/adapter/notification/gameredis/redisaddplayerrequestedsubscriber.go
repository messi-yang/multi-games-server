package gameredis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/presenter/presenterdto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisAddPlayerRequestedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
}

func NewRedisAddPlayerRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) RedisAddPlayerRequestedIntegrationEvent {
	return RedisAddPlayerRequestedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
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
