package coordinatesupdatedeventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/coordinatesupdatedevent"
	"github.com/google/uuid"
)

type coordinatesUpdatedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type coordinatesUpdatedEventCallback = func(coordinateDTOs []coordinatedto.CoordinateDTO)

var coordinatesUpdatedEventInstance *coordinatesUpdatedEventBus

func GetCoordinatesUpdatedEventBus() coordinatesupdatedevent.CoordinatesUpdatedEvent {
	if coordinatesUpdatedEventInstance == nil {
		coordinatesUpdatedEventInstance = &coordinatesUpdatedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "COORDINATES_UPDATED",
		}
	}
	return coordinatesUpdatedEventInstance
}

func (gue *coordinatesUpdatedEventBus) Publish(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinateDTOs)
}

func (gue *coordinatesUpdatedEventBus) Subscribe(gameId uuid.UUID, callback coordinatesUpdatedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
