package gamesocketcontroller

import (
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
	areaUpdatedEventType        eventType = "AREA_UPDATED"
	coordinatesUpdatedEventType eventType = "COORDINATES_UPDATED"
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

type coordinatesUpdatedEventPayload struct {
	Coordinates []coordinatedto.CoordinateDTO `json:"coordinates"`
	Units       []unitdto.UnitDTO             `json:"units"`
}
type coordinatesUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload coordinatesUpdatedEventPayload `json:"payload"`
}

func constructCoordinatesUpdatedEvent(coordinateDTOs []coordinatedto.CoordinateDTO, unitDTOs []unitdto.UnitDTO) *coordinatesUpdatedEvent {
	return &coordinatesUpdatedEvent{
		Type: coordinatesUpdatedEventType,
		Payload: coordinatesUpdatedEventPayload{
			Coordinates: coordinateDTOs,
			Units:       unitDTOs,
		},
	}
}

type areaUpdatedEventPayload struct {
	Area    areadto.AreaDTO       `json:"area"`
	UnitMap unitmapdto.UnitMapDTO `json:"unitMap"`
}
type areaUpdatedEvent struct {
	Type    eventType               `json:"type"`
	Payload areaUpdatedEventPayload `json:"payload"`
}

func constructAreaUpdatedEvent(gameAreaDTO areadto.AreaDTO, unitMap unitmapdto.UnitMapDTO) *areaUpdatedEvent {
	return &areaUpdatedEvent{
		Type: areaUpdatedEventType,
		Payload: areaUpdatedEventPayload{
			Area:    gameAreaDTO,
			UnitMap: unitMap,
		},
	}
}
