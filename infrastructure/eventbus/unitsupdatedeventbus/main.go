package unitsupdatedeventbus

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/unitsupdatedevent"
	"github.com/asaskevich/EventBus"
	"github.com/google/uuid"
)

type unitsUpdatedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type unitsUpdatedEventCallback = func(coordinateDTOs []coordinatedto.CoordinateDTO)

var unitsUpdatedEventInstance *unitsUpdatedEventBus

func GetUnitsUpdatedEventBus() unitsupdatedevent.UnitsUpdatedEvent {
	if unitsUpdatedEventInstance == nil {
		unitsUpdatedEventInstance = &unitsUpdatedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "UNITS_UPDATED",
		}
	}
	return unitsUpdatedEventInstance
}

func (gue *unitsUpdatedEventBus) Publish(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinateDTOs)
}

func (gue *unitsUpdatedEventBus) Subscribe(gameId uuid.UUID, callback unitsUpdatedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
