package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredis"
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

var RedisRemovePlayerRequestedListenerChannel string = "remove-player-requested"

type RedisRemovePlayerRequestedListener struct {
	redisMessageSubscriber *commonredis.RedisMessageSubscriber
}

type redisRedisRemovePlayerRequestedListenerConfiguration func(listener *RedisRemovePlayerRequestedListener) error

func NewRedisRemovePlayerRequestedListener(cfgs ...redisRedisRemovePlayerRequestedListenerConfiguration) (*RedisRemovePlayerRequestedListener, error) {
	t := &RedisRemovePlayerRequestedListener{
		redisMessageSubscriber: commonredis.NewRedisMessageSubscriber(),
	}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (listener *RedisRemovePlayerRequestedListener) Subscribe(subscriber func(RedisRemovePlayerRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisMessageSubscriber.Subscribe(RedisRemovePlayerRequestedListenerChannel, func(message []byte) {
		var event RedisRemovePlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
