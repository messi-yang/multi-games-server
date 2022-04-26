package gamesocketcontroller

import "github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"

type actionType string

const (
	watchGameUnitsActionType actionType = "WATCH_GAME_UNITS"
)

type action struct {
	Type actionType `json:"type"`
}

type watchGameUnitsAction struct {
	Type    actionType `json:"type"`
	Payload struct {
		Area gameservice.GameArea `json:"area"`
	} `json:"payload"`
}
