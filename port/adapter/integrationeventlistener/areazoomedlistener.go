package integrationeventlistener

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type AreaZoomedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  uuid.UUID                          `json:"playerId"`
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewAreaZoomedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) AreaZoomedIntegrationEvent {
	return AreaZoomedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  playerId,
		Area:      area,
		UnitBlock: unitBlock,
	}
}

func AreaZoomedListenerChannel(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", gameId, playerId)
}

type AreaZoomedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisAreaZoomedListenerConfiguration func(listener *AreaZoomedListener) error

func NewAreaZoomedListener(cfgs ...redisAreaZoomedListenerConfiguration) (*AreaZoomedListener, error) {
	t := &AreaZoomedListener{
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

func (listener *AreaZoomedListener) Subscribe(gameId uuid.UUID, playerId uuid.UUID, subscriber func(AreaZoomedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(AreaZoomedListenerChannel(gameId, playerId), func(message []byte) {
		var areaZoomedIntegrationEvent AreaZoomedIntegrationEvent
		json.Unmarshal(message, &areaZoomedIntegrationEvent)
		subscriber(areaZoomedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
