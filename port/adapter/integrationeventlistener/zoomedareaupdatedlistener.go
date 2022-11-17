package integrationeventlistener

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type ZoomedAreaUpdatedIntegrationEvent struct {
	GameId    uuid.UUID                          `json:"gameId"`
	PlayerId  uuid.UUID                          `json:"playerId"`
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) ZoomedAreaUpdatedIntegrationEvent {
	return ZoomedAreaUpdatedIntegrationEvent{
		GameId:    gameId,
		PlayerId:  playerId,
		Area:      area,
		UnitBlock: unitBlock,
	}
}

func ZoomedAreaUpdatedListenerChannel(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", gameId, playerId)
}

type ZoomedAreaUpdatedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisZoomedAreaUpdatedListenerConfiguration func(listener *ZoomedAreaUpdatedListener) error

func NewZoomedAreaUpdatedListener(cfgs ...redisZoomedAreaUpdatedListenerConfiguration) (*ZoomedAreaUpdatedListener, error) {
	t := &ZoomedAreaUpdatedListener{
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

func (listener *ZoomedAreaUpdatedListener) Subscribe(gameId uuid.UUID, playerId uuid.UUID, subscriber func(ZoomedAreaUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(ZoomedAreaUpdatedListenerChannel(gameId, playerId), func(message []byte) {
		var zoomedAreaUpdatedIntegrationEvent ZoomedAreaUpdatedIntegrationEvent
		json.Unmarshal(message, &zoomedAreaUpdatedIntegrationEvent)
		subscriber(zoomedAreaUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
