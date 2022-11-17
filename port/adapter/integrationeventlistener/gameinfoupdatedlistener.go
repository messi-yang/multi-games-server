package integrationeventlistener

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type GameInfoUpdatedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  uuid.UUID                          `json:"playerId"`
	Dimension presenterdto.DimensionPresenterDto `json:"dimension"`
}

func NewGameInfoUpdatedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, dimensionPresenterDto presenterdto.DimensionPresenterDto) GameInfoUpdatedIntegrationEvent {
	return GameInfoUpdatedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  playerId,
		Dimension: dimensionPresenterDto,
	}
}

func GameInfoUpdatedListenerChannel(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", gameId, playerId)
}

type GameInfoUpdatedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisGameInfoUpdatedListenerConfiguration func(listener *GameInfoUpdatedListener) error

func NewGameInfoUpdatedListener(cfgs ...redisGameInfoUpdatedListenerConfiguration) (*GameInfoUpdatedListener, error) {
	t := &GameInfoUpdatedListener{
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

func (listener *GameInfoUpdatedListener) Subscribe(gameId uuid.UUID, playerId uuid.UUID, subscriber func(GameInfoUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(GameInfoUpdatedListenerChannel(gameId, playerId), func(message []byte) {
		var gameInfoUpdatedIntegrationEvent GameInfoUpdatedIntegrationEvent
		json.Unmarshal(message, &gameInfoUpdatedIntegrationEvent)
		subscriber(gameInfoUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
