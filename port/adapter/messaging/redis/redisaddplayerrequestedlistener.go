package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/google/uuid"
)

type RedisAddPlayerRequestedIntegrationEvent struct {
	GameId   uuid.UUID `json:"gameId"`
	PlayerId uuid.UUID `json:"playerId"`
}

func NewRedisAddPlayerRequestedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID) RedisAddPlayerRequestedIntegrationEvent {
	return RedisAddPlayerRequestedIntegrationEvent{
		GameId:   gameId,
		PlayerId: playerId,
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
