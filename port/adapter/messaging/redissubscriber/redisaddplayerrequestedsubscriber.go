package redissubscriber

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredissubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
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
	redisMessageSubscriber *commonredissubscriber.RedisMessageSubscriber
}

type redisRedisAddPlayerRequestedSubscriberConfiguration func(subscriber *RedisAddPlayerRequestedSubscriber) error

func NewRedisAddPlayerRequestedSubscriber(cfgs ...redisRedisAddPlayerRequestedSubscriberConfiguration) (commonredissubscriber.RedisSubscriber[RedisAddPlayerRequestedIntegrationEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisMessageSubscriber: commonredissubscriber.NewRedisMessageSubscriber(),
	}, nil
}

func (subscriber *RedisAddPlayerRequestedSubscriber) Subscribe(handler func(RedisAddPlayerRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisMessageSubscriber.Subscribe(RedisAddPlayerRequestedSubscriberChannel, func(message []byte) {
		var event RedisAddPlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
