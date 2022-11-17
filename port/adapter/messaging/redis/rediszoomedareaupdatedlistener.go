package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type RedisZoomedAreaUpdatedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  uuid.UUID                          `json:"playerId"`
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewRedisZoomedAreaUpdatedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) RedisZoomedAreaUpdatedIntegrationEvent {
	return RedisZoomedAreaUpdatedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  playerId,
		Area:      area,
		UnitBlock: unitBlock,
	}
}

func RedisZoomedAreaUpdatedListenerChannel(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", gameId, playerId)
}

type RedisZoomedAreaUpdatedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisZoomedAreaUpdatedListenerConfiguration func(listener *RedisZoomedAreaUpdatedListener) error

func NewRedisZoomedAreaUpdatedListener(cfgs ...redisRedisZoomedAreaUpdatedListenerConfiguration) (*RedisZoomedAreaUpdatedListener, error) {
	t := &RedisZoomedAreaUpdatedListener{
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

func (listener *RedisZoomedAreaUpdatedListener) Subscribe(gameId uuid.UUID, playerId uuid.UUID, subscriber func(RedisZoomedAreaUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisZoomedAreaUpdatedListenerChannel(gameId, playerId), func(message []byte) {
		var redisZoomedAreaUpdatedIntegrationEvent RedisZoomedAreaUpdatedIntegrationEvent
		json.Unmarshal(message, &redisZoomedAreaUpdatedIntegrationEvent)
		subscriber(redisZoomedAreaUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
