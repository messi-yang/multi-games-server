package gamesocketappservice

import (
	"fmt"
)

func CreateGamePlayerChannel(gameIdDto string, playerIdDto string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s", gameIdDto, playerIdDto)
}

type GameSocketIntEventName string

const (
	PlayersUpdatedGameSocketIntEventName GameSocketIntEventName = "PLAYERS_UPDATED"
	ViewUpdatedGameSocketIntEventName    GameSocketIntEventName = "VIEW_UPDATED"
)

type GameSocketIntEvent struct {
	Name GameSocketIntEventName `json:"name"`
}

type PlayersUpdatedIntEvent struct {
	Name   GameSocketIntEventName `json:"name"`
	GameId string                 `json:"gameId"`
}

func NewPlayersUpdatedIntEvent(gameIdDto string) PlayersUpdatedIntEvent {
	return PlayersUpdatedIntEvent{
		Name:   PlayersUpdatedGameSocketIntEventName,
		GameId: gameIdDto,
	}
}

type ViewUpdatedIntEvent struct {
	Name   GameSocketIntEventName `json:"name"`
	GameId string                 `json:"gameId"`
}

func NewViewUpdatedIntEvent(gameIdDto string) ViewUpdatedIntEvent {
	return ViewUpdatedIntEvent{
		Name:   ViewUpdatedGameSocketIntEventName,
		GameId: gameIdDto,
	}
}
