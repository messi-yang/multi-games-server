package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/areadto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/gameunitdto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
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
			MapSize: mapsizedto.MapSizeDTO{
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

func constructAreaUpdatedEvent(gameArea *valueobject.Area, gameUnits *[][]gameunitdto.GameUnitDTO) *areaUpdatedEvent {
	return &areaUpdatedEvent{
		Type: areaUpdatedEventType,
		Payload: areaUpdatedEventPayload{
			Area: areadto.AreaDTO{
				From: coordinatedto.CoordinateDTO{
					X: gameArea.GetFrom().GetX(),
					Y: gameArea.GetFrom().GetY(),
				},
				To: coordinatedto.CoordinateDTO{
					X: gameArea.GetTo().GetX(),
					Y: gameArea.GetTo().GetY(),
				},
			},
			Units: *gameUnits,
		},
	}
}
