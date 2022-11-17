package integrationeventlistener

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/google/uuid"
)

type RemovePlayerRequestedIntegrationEvent struct {
	GameId   uuid.UUID `json:"gameId"`
	PlayerId uuid.UUID `json:"playerId"`
}

func NewRemovePlayerRequestedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID) RemovePlayerRequestedIntegrationEvent {
	return RemovePlayerRequestedIntegrationEvent{
		GameId:   gameId,
		PlayerId: playerId,
	}
}

var RemovePlayerRequestedListenerChannel string = "remove-player-requested"

type RemovePlayerRequestedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRemovePlayerRequestedListenerConfiguration func(listener *RemovePlayerRequestedListener) error

func NewRemovePlayerRequestedListener(cfgs ...redisRemovePlayerRequestedListenerConfiguration) (*RemovePlayerRequestedListener, error) {
	t := &RemovePlayerRequestedListener{
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

func (listener *RemovePlayerRequestedListener) Subscribe(subscriber func(RemovePlayerRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RemovePlayerRequestedListenerChannel, func(message []byte) {
		var removePlayerRequestedIntegrationEvent RemovePlayerRequestedIntegrationEvent
		json.Unmarshal(message, &removePlayerRequestedIntegrationEvent)
		subscriber(removePlayerRequestedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
