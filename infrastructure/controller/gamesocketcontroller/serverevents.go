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
	unitMapUpdatedEventType     eventType = "UNIT_MAP_UPDATED"
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
	MapSize mapsizedto.MapSizeDTO `json:"mapSize"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

func constructInformationUpdatedEvent(mapSizeDTO mapsizedto.MapSizeDTO) *informationUpdatedEvent {
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize: mapSizeDTO,
		},
	}
}

type unitsRevivedEventPayload struct {
	Coordinates []coordinatedto.CoordinateDTO `json:"coordinates"`
	Units       []unitdto.UnitDTO             `json:"units"`
	UpdatedAt   time.Time                     `json:"updateAt"`
}
type unitsRevivedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsRevivedEventPayload `json:"payload"`
}

func constructUnitsRevivedEvent(coordinateDTOs []coordinatedto.CoordinateDTO, unitDTOs []unitdto.UnitDTO, updatedAt time.Time) *unitsRevivedEvent {
	return &unitsRevivedEvent{
		Type: unitsRevivedEventType,
		Payload: unitsRevivedEventPayload{
			Coordinates: coordinateDTOs,
			Units:       unitDTOs,
			UpdatedAt:   updatedAt,
		},
	}
}

type unitMapUpdatedEventPayload struct {
	Area      areadto.AreaDTO       `json:"area"`
	UnitMap   unitmapdto.UnitMapDTO `json:"unitMap"`
	UpdatedAt time.Time             `json:"updateAt"`
}
type unitMapUpdatedEvent struct {
	Type    eventType                  `json:"type"`
	Payload unitMapUpdatedEventPayload `json:"payload"`
}

func constructUnitMapUpdated(gameAreaDTO areadto.AreaDTO, unitMap unitmapdto.UnitMapDTO, updatedAt time.Time) *unitMapUpdatedEvent {
	return &unitMapUpdatedEvent{
		Type: unitMapUpdatedEventType,
		Payload: unitMapUpdatedEventPayload{
			Area:      gameAreaDTO,
			UnitMap:   unitMap,
			UpdatedAt: updatedAt,
		},
	}
}

type unitMapReceivedEventPayload struct {
	Area       areadto.AreaDTO       `json:"area"`
	UnitMap    unitmapdto.UnitMapDTO `json:"unitMap"`
	ReceivedAt time.Time             `json:"receivedAt"`
}
type unitMapReceivedEvent struct {
	Type    eventType                   `json:"type"`
	Payload unitMapReceivedEventPayload `json:"payload"`
}

func constructUnitMapReceived(gameAreaDTO areadto.AreaDTO, unitMap unitmapdto.UnitMapDTO, receivedAt time.Time) *unitMapReceivedEvent {
	return &unitMapReceivedEvent{
		Type: unitMapReceivedEventType,
		Payload: unitMapReceivedEventPayload{
			Area:       gameAreaDTO,
			UnitMap:    unitMap,
			ReceivedAt: receivedAt,
		},
	}
}
