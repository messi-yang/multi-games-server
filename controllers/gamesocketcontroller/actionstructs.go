package gamesocketcontroller

import "github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"

type actionType string

const (
	watchUnitsActionType actionType = "WATCH_UNITS"
)

type action struct {
	Type actionType `json:"type"`
}

type watchUnitsActionPayload struct {
	Area gameservice.GameArea `json:"area"`
}
type watchUnitsAction struct {
	Type    actionType              `json:"type"`
	Payload watchUnitsActionPayload `json:"payload"`
}
