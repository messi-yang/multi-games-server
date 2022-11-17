package integrationeventlistener

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type ReviveUnitsRequestedIntegrationEvent struct {
	GameId      uuid.UUID                             `json:"gameId"`
	Coordinates []presenterdto.CoordinatePresenterDto `json:"coordinates"`
}

func NewReviveUnitsRequestedIntegrationEvent(gameId uuid.UUID, coordinatePresenterDtos []presenterdto.CoordinatePresenterDto) ReviveUnitsRequestedIntegrationEvent {
	return ReviveUnitsRequestedIntegrationEvent{
		GameId:      gameId,
		Coordinates: coordinatePresenterDtos,
	}
}

var ReviveUnitsRequestedListenerChannel string = "revive-units-requested"

type ReviveUnitsRequestedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisReviveUnitsRequestedListenerConfiguration func(listener *ReviveUnitsRequestedListener) error

func NewReviveUnitsRequestedListener(cfgs ...redisReviveUnitsRequestedListenerConfiguration) (*ReviveUnitsRequestedListener, error) {
	t := &ReviveUnitsRequestedListener{
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

func (listener *ReviveUnitsRequestedListener) Subscribe(subscriber func(ReviveUnitsRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(ReviveUnitsRequestedListenerChannel, func(message []byte) {
		var reviveUnitsRequestedIntegrationEvent ReviveUnitsRequestedIntegrationEvent
		json.Unmarshal(message, &reviveUnitsRequestedIntegrationEvent)
		subscriber(reviveUnitsRequestedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
