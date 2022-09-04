package gamesocketcontroller

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitmapdto"
)

type eventType string

const (
	errorHappenedEventType      eventType = "ERROR"
	informationUpdatedEventType eventType = "INFORMATION_UPDATED"
	areaZoomedEventType         eventType = "AREA_ZOOMED"
	unitMapTickedEventType      eventType = "UNIT_MAP_TICKED"
	unitsRevivedEventType       eventType = "UNITS_REVIVED"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

func constructErrorHappenedEvent(clientMessage string) *errorHappenedEvent {
	return &errorHappenedEvent{
		Type: errorHappenedEventType,
		Payload: errorHappenedEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

type informationUpdatedEventPayload struct {
	MapSize mapsizedto.Dto `json:"mapSize"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

func constructInformationUpdatedEvent(mapSizeDto mapsizedto.Dto) *informationUpdatedEvent {
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize: mapSizeDto,
		},
	}
}

type unitsRevivedEventPayload struct {
	Coordinates []coordinatedto.Dto `json:"coordinates"`
	Units       []unitdto.Dto       `json:"units"`
	UpdatedAt   time.Time           `json:"updateAt"`
}
type unitsRevivedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsRevivedEventPayload `json:"payload"`
}

func constructUnitsRevivedEvent(coordinateDtos []coordinatedto.Dto, unitDtos []unitdto.Dto, updatedAt time.Time) *unitsRevivedEvent {
	return &unitsRevivedEvent{
		Type: unitsRevivedEventType,
		Payload: unitsRevivedEventPayload{
			Coordinates: coordinateDtos,
			Units:       unitDtos,
			UpdatedAt:   updatedAt,
		},
	}
}

type unitMapTickedEventPayload struct {
	Area      areadto.Dto    `json:"area"`
	UnitMap   unitmapdto.Dto `json:"unitMap"`
	UpdatedAt time.Time      `json:"updateAt"`
}
type unitMapTickedEvent struct {
	Type    eventType                 `json:"type"`
	Payload unitMapTickedEventPayload `json:"payload"`
}

func constructUnitMapTicked(gameDto areadto.Dto, unitMap unitmapdto.Dto, updatedAt time.Time) *unitMapTickedEvent {
	return &unitMapTickedEvent{
		Type: unitMapTickedEventType,
		Payload: unitMapTickedEventPayload{
			Area:      gameDto,
			UnitMap:   unitMap,
			UpdatedAt: updatedAt,
		},
	}
}

type areaZoomedEventPayload struct {
	Area    areadto.Dto    `json:"area"`
	UnitMap unitmapdto.Dto `json:"unitMap"`
}
type areaZoomedEvent struct {
	Type    eventType              `json:"type"`
	Payload areaZoomedEventPayload `json:"payload"`
}

func constructAreaZoomedEvent(gameDto areadto.Dto, unitMap unitmapdto.Dto) *areaZoomedEvent {
	return &areaZoomedEvent{
		Type: areaZoomedEventType,
		Payload: areaZoomedEventPayload{
			Area:    gameDto,
			UnitMap: unitMap,
		},
	}
}
