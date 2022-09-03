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
	unitMapReceivedEventType    eventType = "UNIT_MAP_RECEIVED"
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
	MapSize mapsizedto.DTO `json:"mapSize"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

func constructInformationUpdatedEvent(mapSizeDTO mapsizedto.DTO) *informationUpdatedEvent {
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize: mapSizeDTO,
		},
	}
}

type unitsRevivedEventPayload struct {
	Coordinates []coordinatedto.DTO `json:"coordinates"`
	Units       []unitdto.DTO       `json:"units"`
	UpdatedAt   time.Time           `json:"updateAt"`
}
type unitsRevivedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsRevivedEventPayload `json:"payload"`
}

func constructUnitsRevivedEvent(coordinateDTOs []coordinatedto.DTO, unitDTOs []unitdto.DTO, updatedAt time.Time) *unitsRevivedEvent {
	return &unitsRevivedEvent{
		Type: unitsRevivedEventType,
		Payload: unitsRevivedEventPayload{
			Coordinates: coordinateDTOs,
			Units:       unitDTOs,
			UpdatedAt:   updatedAt,
		},
	}
}

type unitMapTickedEventPayload struct {
	Area      areadto.DTO    `json:"area"`
	UnitMap   unitmapdto.DTO `json:"unitMap"`
	UpdatedAt time.Time      `json:"updateAt"`
}
type unitMapTickedEvent struct {
	Type    eventType                 `json:"type"`
	Payload unitMapTickedEventPayload `json:"payload"`
}

func constructUnitMapTicked(gameDTO areadto.DTO, unitMap unitmapdto.DTO, updatedAt time.Time) *unitMapTickedEvent {
	return &unitMapTickedEvent{
		Type: unitMapTickedEventType,
		Payload: unitMapTickedEventPayload{
			Area:      gameDTO,
			UnitMap:   unitMap,
			UpdatedAt: updatedAt,
		},
	}
}

type unitMapReceivedEventPayload struct {
	Area       areadto.DTO    `json:"area"`
	UnitMap    unitmapdto.DTO `json:"unitMap"`
	ReceivedAt time.Time      `json:"receivedAt"`
}
type unitMapReceivedEvent struct {
	Type    eventType                   `json:"type"`
	Payload unitMapReceivedEventPayload `json:"payload"`
}

func constructUnitMapReceived(gameDTO areadto.DTO, unitMap unitmapdto.DTO, receivedAt time.Time) *unitMapReceivedEvent {
	return &unitMapReceivedEvent{
		Type: unitMapReceivedEventType,
		Payload: unitMapReceivedEventPayload{
			Area:       gameDTO,
			UnitMap:    unitMap,
			ReceivedAt: receivedAt,
		},
	}
}
