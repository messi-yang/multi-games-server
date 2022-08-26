package gameunitsupdatedeventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitsupdatedevent"
	"github.com/google/uuid"
)

type gameUnitsUpdatedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameUnitsUpdatedEventCallback = func(coordinateDTOs []coordinatedto.CoordinateDTO)

var gameUnitsUpdatedEventInstance *gameUnitsUpdatedEventBus

func GetCoordinatesUpdatedEventBus() gameunitsupdatedevent.CoordinatesUpdatedEvent {
	if gameUnitsUpdatedEventInstance == nil {
		gameUnitsUpdatedEventInstance = &gameUnitsUpdatedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "COORDINATES_UPDATED",
		}
	}
	return gameUnitsUpdatedEventInstance
}

func (gue *gameUnitsUpdatedEventBus) Publish(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinateDTOs)
}

func (gue *gameUnitsUpdatedEventBus) Subscribe(gameId uuid.UUID, callback gameUnitsUpdatedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
