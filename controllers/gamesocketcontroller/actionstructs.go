package gamesocketcontroller

import "github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"

type actionType string

const (
	watchGameBlock actionType = "WATCH_GAME_BLOCK"
)

type action struct {
	Type actionType `json:"type"`
}

type watchGameBlockAction struct {
	Type    actionType `json:"type"`
	Payload struct {
		Area gameservice.GameArea `json:"area"`
	} `json:"payload"`
}
