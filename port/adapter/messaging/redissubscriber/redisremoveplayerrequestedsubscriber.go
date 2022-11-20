package redissubscriber

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredissubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
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
	redisMessageSubscriber *commonredissubscriber.RedisMessageSubscriber
}

func NewRedisRemovePlayerRequestedSubscriber() (commonredissubscriber.RedisSubscriber[RedisRemovePlayerRequestedIntegrationEvent], error) {
	return &RedisRemovePlayerRequestedSubscriber{
		redisMessageSubscriber: commonredissubscriber.NewRedisMessageSubscriber(),
	}, nil
}

func (subscriber *RedisRemovePlayerRequestedSubscriber) Subscribe(handler func(RedisRemovePlayerRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisMessageSubscriber.Subscribe(RedisRemovePlayerRequestedSubscriberChannel, func(message []byte) {
		var event RedisRemovePlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
