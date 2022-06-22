package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/areadto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/gameunitdto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/mapsizedto"
)

type eventType string

const (
	errorHappenedEventType      eventType = "ERROR"
	informationUpdatedEventType           = "INFORMATION_UPDATED"
	areaUpdatedEventType                  = "AREA_UPDATED"
	unitsUpdatedEventType                 = "UNITS_UPDATED"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

type informationUpdatedEventPayload struct {
	MapSize      mapsizedto.MapSizeDTO `json:"mapSize"`
	PlayersCount int                   `json:"playersCount"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

type unitsUpdatedEventPayloadItem struct {
	Coordinate coordinatedto.CoordinateDTO `json:"coordinate"`
	Unit       gameunitdto.GameUnitDTO     `json:"unit"`
}

type unitsUpdatedEventPayload struct {
	Items []unitsUpdatedEventPayloadItem `json:"items"`
}
type unitsUpdatedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsUpdatedEventPayload `json:"payload"`
}

type areaUpdatedEventPayload struct {
	Area  areadto.AreaDTO             `json:"area"`
	Units [][]gameunitdto.GameUnitDTO `json:"units"`
}
type areaUpdatedEvent struct {
	Type    eventType               `json:"type"`
	Payload areaUpdatedEventPayload `json:"payload"`
}
