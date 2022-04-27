package gamesocketcontroller

import "github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"

func constructErrorHappenedEvent(clientMessage string) *errorHappenedEvent {
	return &errorHappenedEvent{
		Type: errorHappenedEventType,
		Payload: errorHappenedEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

func constructUnitsUpdatedEvent(gameArea *gameservice.GameArea, gameUnits *[][]*gameservice.GameUnit) *unitsUpdatedEvent {
	return &unitsUpdatedEvent{
		Type: unitsUpdatedEventType,
		Payload: unitsUpdatedEventPayload{
			Area:  *gameArea,
			Units: *gameUnits,
		},
	}
}
