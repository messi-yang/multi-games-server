package integrationeventlistener

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/google/uuid"
)

type AddPlayerRequestedIntegrationEvent struct {
	GameId   uuid.UUID `json:"gameId"`
	PlayerId uuid.UUID `json:"playerId"`
}

func NewAddPlayerRequestedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID) AddPlayerRequestedIntegrationEvent {
	return AddPlayerRequestedIntegrationEvent{
		GameId:   gameId,
		PlayerId: playerId,
	}
}

var AddPlayerRequestedListenerChannel string = "add-player-requested"

type AddPlayerRequestedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisAddPlayerRequestedListenerConfiguration func(listener *AddPlayerRequestedListener) error

func NewAddPlayerRequestedListener(cfgs ...redisAddPlayerRequestedListenerConfiguration) (*AddPlayerRequestedListener, error) {
	t := &AddPlayerRequestedListener{
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

func (listener *AddPlayerRequestedListener) Subscribe(subscriber func(AddPlayerRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(AddPlayerRequestedListenerChannel, func(message []byte) {
		var addPlayerRequestedIntegrationEvent AddPlayerRequestedIntegrationEvent
		json.Unmarshal(message, &addPlayerRequestedIntegrationEvent)
		subscriber(addPlayerRequestedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
