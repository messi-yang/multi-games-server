package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

type eventType string

const (
	errorHappenedEventType      eventType = "ERROR"
	informationUpdatedEventType           = "INFORMATION_UPDATED"
	areaUpdatedEventType                  = "AREA_UPDATED"
	unitsUpdatedEventType                 = "UNITS_UPDATED"
	playerJoinedEventType                 = "PLAYER_JOINED"
	playerLeftEventType                   = "PLAYER_LEFT"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

type informationUpdatedEventPayload struct {
	MapSize      valueobject.GameSize `json:"mapSize"`
	PlayersCount int                  `json:"playersCount"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

type unitsUpdatedEventPayloadItem struct {
	Coordinate gameservice.GameCoordinate `json:"coordinate"`
	Unit       valueobject.GameUnit       `json:"unit"`
}

type unitsUpdatedEventPayload struct {
	Items []unitsUpdatedEventPayloadItem `json:"items"`
}
type unitsUpdatedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsUpdatedEventPayload `json:"payload"`
}

type areaUpdatedEventPayload struct {
	Area  gameservice.GameArea      `json:"area"`
	Units [][]*valueobject.GameUnit `json:"units"`
}
type areaUpdatedEvent struct {
	Type    eventType               `json:"type"`
	Payload areaUpdatedEventPayload `json:"payload"`
}

type playerJoinedEventPayload any
type playerJoinedEvent struct {
	Type    eventType                `json:"type"`
	Payload playerJoinedEventPayload `json:"payload"`
}

type playerLeftEventPayload any
type playerLeftEvent struct {
	Type    eventType              `json:"type"`
	Payload playerLeftEventPayload `json:"payload"`
}
