package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
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
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisAddPlayerRequestedListenerConfiguration func(listener *RedisAddPlayerRequestedListener) error

func NewRedisAddPlayerRequestedListener(cfgs ...redisRedisAddPlayerRequestedListenerConfiguration) (*RedisAddPlayerRequestedListener, error) {
	t := &RedisAddPlayerRequestedListener{
		redisInfrastructureService: service.NewRedisInfrastructureService(),
	}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (listener *RedisAddPlayerRequestedListener) Subscribe(subscriber func(RedisAddPlayerRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisAddPlayerRequestedListenerChannel, func(message []byte) {
		var event RedisAddPlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
