package gamesocketcontroller

import "github.com/DumDumGeniuss/game-of-liberty-computer/entities/gameentity"

type coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type area struct {
	From coordinate `json:"from"`
	To   coordinate `json:"to"`
}

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
		Area area `json:"area"`
	} `json:"payload"`
}

type eventType string

const (
	gaemBlockUpdated eventType = "GAME_BLOCK_UPDATED"
)

type gameBlockUpdatedEventPayload struct {
	Area  area                     `json:"area"`
	Units [][]*gameentity.GameUnit `json:"units"`
}

type gameBlockUpdatedEvent struct {
	Type    eventType                    `json:"type"`
	Payload gameBlockUpdatedEventPayload `json:"payload"`
}
