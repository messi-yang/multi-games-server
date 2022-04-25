package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

type eventType string

const (
	gaemBlockUpdated eventType = "GAME_BLOCK_UPDATED"
)

type gameBlockUpdatedEvent struct {
	Type    eventType `json:"type"`
	Payload struct {
		Area  gameservice.GameArea      `json:"area"`
		Units [][]*gameservice.GameUnit `json:"units"`
	} `json:"payload"`
}
