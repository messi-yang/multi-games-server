package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/google/uuid"
)

type RedisRemovePlayerRequestedIntegrationEvent struct {
	GameId   uuid.UUID `json:"gameId"`
	PlayerId uuid.UUID `json:"playerId"`
}

func NewRedisRemovePlayerRequestedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID) RedisRemovePlayerRequestedIntegrationEvent {
	return RedisRemovePlayerRequestedIntegrationEvent{
		GameId:   gameId,
		PlayerId: playerId,
	}
}

var RedisRemovePlayerRequestedListenerChannel string = "remove-player-requested"

type RedisRemovePlayerRequestedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisRemovePlayerRequestedListenerConfiguration func(listener *RedisRemovePlayerRequestedListener) error

func NewRedisRemovePlayerRequestedListener(cfgs ...redisRedisRemovePlayerRequestedListenerConfiguration) (*RedisRemovePlayerRequestedListener, error) {
	t := &RedisRemovePlayerRequestedListener{
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

func (listener *RedisRemovePlayerRequestedListener) Subscribe(subscriber func(RedisRemovePlayerRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisRemovePlayerRequestedListenerChannel, func(message []byte) {
		var event RedisRemovePlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
