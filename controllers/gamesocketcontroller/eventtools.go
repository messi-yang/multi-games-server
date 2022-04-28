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

func constructGameInfoUpdatedEvent(mapSize *gameservice.GameSize, playersCount int) *gameInfoUpdatedEvent {
	return &gameInfoUpdatedEvent{
		Type: gameInfoUpdatedEventType,
		Payload: gameInfoUpdatedEventPayload{
			MapSize:      *mapSize,
			PlayersCount: playersCount,
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

func constructPlayerJoinedEvent() *playerJoinedEvent {
	return &playerJoinedEvent{
		Type:    playerJoinedEventType,
		Payload: nil,
	}
}

func constructPlayerLeftEvent() *playerLeftEvent {
	return &playerLeftEvent{
		Type:    playerLeftEventType,
		Payload: nil,
	}
}
