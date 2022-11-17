package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type RedisGameInfoUpdatedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  uuid.UUID                          `json:"playerId"`
	Dimension presenterdto.DimensionPresenterDto `json:"dimension"`
}

func NewRedisGameInfoUpdatedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, dimensionPresenterDto presenterdto.DimensionPresenterDto) RedisGameInfoUpdatedIntegrationEvent {
	return RedisGameInfoUpdatedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  playerId,
		Dimension: dimensionPresenterDto,
	}
}

func RedisGameInfoUpdatedListenerChannel(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", gameId, playerId)
}

type RedisGameInfoUpdatedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisGameInfoUpdatedListenerConfiguration func(listener *RedisGameInfoUpdatedListener) error

func NewRedisGameInfoUpdatedListener(cfgs ...redisRedisGameInfoUpdatedListenerConfiguration) (*RedisGameInfoUpdatedListener, error) {
	t := &RedisGameInfoUpdatedListener{
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

func (listener *RedisGameInfoUpdatedListener) Subscribe(gameId uuid.UUID, playerId uuid.UUID, subscriber func(RedisGameInfoUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisGameInfoUpdatedListenerChannel(gameId, playerId), func(message []byte) {
		var event RedisGameInfoUpdatedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
