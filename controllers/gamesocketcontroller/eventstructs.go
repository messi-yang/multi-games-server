package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

type eventType string

const (
	errorHappenedEventType    eventType = "ERROR"
	gaemUnitsUpdatedEventType           = "GAME_UNITS_UPDATED"
)

type errorHappenedEvent struct {
	Type    eventType `json:"type"`
	Payload struct {
		ClientMessage string `json:"clientMessage"`
	} `json:"payload"`
}

type gameUnitsUpdatedEvent struct {
	Type    eventType `json:"type"`
	Payload struct {
		Area  gameservice.GameArea      `json:"area"`
		Units [][]*gameservice.GameUnit `json:"units"`
	} `json:"payload"`
}
