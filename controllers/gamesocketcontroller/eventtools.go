package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

func constructErrorHappenedEvent(clientMessage string) *errorHappenedEvent {
	return &errorHappenedEvent{
		Type: errorHappenedEventType,
		Payload: errorHappenedEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

func constructInformationUpdatedEvent(mapSize *valueobject.MapSize, playersCount int) *informationUpdatedEvent {
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize: MapSizeDTO{
				Width:  mapSize.GetWidth(),
				Height: mapSize.GetHeight(),
			},
			PlayersCount: playersCount,
		},
	}
}

func constructUnitsUpdatedEvent(items *[]unitsUpdatedEventPayloadItem) *unitsUpdatedEvent {
	return &unitsUpdatedEvent{
		Type: unitsUpdatedEventType,
		Payload: unitsUpdatedEventPayload{
			Items: *items,
		},
	}
}

func constructAreaUpdatedEvent(gameArea *gameservice.GameArea, gameUnits *[][]*valueobject.GameUnit) *areaUpdatedEvent {
	return &areaUpdatedEvent{
		Type: areaUpdatedEventType,
		Payload: areaUpdatedEventPayload{
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
