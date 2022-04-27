package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

type eventType string

const (
	errorHappenedEventType eventType = "ERROR"
	unitsUpdatedEventType            = "UNITS_UPDATED"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

type unitsUpdatedEventPayload struct {
	Area  gameservice.GameArea      `json:"area"`
	Units [][]*gameservice.GameUnit `json:"units"`
}
type unitsUpdatedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsUpdatedEventPayload `json:"payload"`
}
