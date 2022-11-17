package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type RedisZoomedAreaUpdatedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  presenterdto.PlayerIdPresenterDto  `json:"playerId"`
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewRedisZoomedAreaUpdatedIntegrationEvent(gameId uuid.UUID, playerId gamecommonmodel.PlayerId, area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) RedisZoomedAreaUpdatedIntegrationEvent {
	return RedisZoomedAreaUpdatedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  presenterdto.NewPlayerIdPresenterDto(playerId),
		Area:      area,
		UnitBlock: unitBlock,
	}
}

func RedisZoomedAreaUpdatedListenerChannel(gameId uuid.UUID, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", gameId, playerId.GetId().String())
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

func (listener *RedisZoomedAreaUpdatedListener) Subscribe(gameId uuid.UUID, playerId gamecommonmodel.PlayerId, subscriber func(RedisZoomedAreaUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisZoomedAreaUpdatedListenerChannel(gameId, playerId), func(message []byte) {
		var redisZoomedAreaUpdatedIntegrationEvent RedisZoomedAreaUpdatedIntegrationEvent
		json.Unmarshal(message, &redisZoomedAreaUpdatedIntegrationEvent)
		subscriber(redisZoomedAreaUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
