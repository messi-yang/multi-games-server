package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type RedisAreaZoomedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  uuid.UUID                          `json:"playerId"`
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewRedisAreaZoomedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) RedisAreaZoomedIntegrationEvent {
	return RedisAreaZoomedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  playerId,
		Area:      area,
		UnitBlock: unitBlock,
	}
}

func RedisAreaZoomedListenerChannel(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", gameId, playerId)
}

type RedisAreaZoomedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisAreaZoomedListenerConfiguration func(listener *RedisAreaZoomedListener) error

func NewRedisAreaZoomedListener(cfgs ...redisRedisAreaZoomedListenerConfiguration) (*RedisAreaZoomedListener, error) {
	t := &RedisAreaZoomedListener{
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

func (listener *RedisAreaZoomedListener) Subscribe(gameId uuid.UUID, playerId uuid.UUID, subscriber func(RedisAreaZoomedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisAreaZoomedListenerChannel(gameId, playerId), func(message []byte) {
		var event RedisAreaZoomedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
