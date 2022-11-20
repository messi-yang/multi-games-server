package redislistener

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredislistener"
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

var RedisAddPlayerRequestedListenerChannel string = "add-player-requested"

type RedisAddPlayerRequestedListener struct {
	redisMessageSubscriber *commonredislistener.RedisMessageSubscriber
}

type redisRedisAddPlayerRequestedListenerConfiguration func(listener *RedisAddPlayerRequestedListener) error

func NewRedisAddPlayerRequestedListener(cfgs ...redisRedisAddPlayerRequestedListenerConfiguration) (commonredislistener.RedisListener[RedisAddPlayerRequestedIntegrationEvent], error) {
	return &RedisAddPlayerRequestedListener{
		redisMessageSubscriber: commonredislistener.NewRedisMessageSubscriber(),
	}, nil
}

func (listener *RedisAddPlayerRequestedListener) Subscribe(subscriber func(RedisAddPlayerRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisMessageSubscriber.Subscribe(RedisAddPlayerRequestedListenerChannel, func(message []byte) {
		var event RedisAddPlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
