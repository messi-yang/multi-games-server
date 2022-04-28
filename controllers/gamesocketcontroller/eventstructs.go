package gamesocketcontroller

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
)

type eventType string

const (
	errorHappenedEventType   eventType = "ERROR"
	gameInfoUpdatedEventType           = "GAME_INFO_UPDATED"
	unitsUpdatedEventType              = "UNITS_UPDATED"
	playerJoinedEventType              = "PLAYER_JOINED"
	playerLeftEventType                = "PLAYER_LEFT"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

type gameInfoUpdatedEventPayload struct {
	MapSize      gameservice.GameSize `json:"mapSize"`
	PlayersCount int                  `json:"playersCount"`
}
type gameInfoUpdatedEvent struct {
	Type    eventType                   `json:"type"`
	Payload gameInfoUpdatedEventPayload `json:"payload"`
}

type unitsUpdatedEventPayload struct {
	Area  gameservice.GameArea      `json:"area"`
	Units [][]*gameservice.GameUnit `json:"units"`
}
type unitsUpdatedEvent struct {
	Type    eventType                `json:"type"`
	Payload unitsUpdatedEventPayload `json:"payload"`
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
