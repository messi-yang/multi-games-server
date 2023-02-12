package gamesocketappservice

import (
	"fmt"
)

func CreateGamePlayerChannel(gameIdVm string, playerIdVm string) string {
	return fmt.Sprintf("GAME_%s_PLAYER_%s", gameIdVm, playerIdVm)
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

func NewPlayersUpdatedIntEvent(gameIdVm string) PlayersUpdatedIntEvent {
	return PlayersUpdatedIntEvent{
		Name:   PlayersUpdatedGameSocketIntEventName,
		GameId: gameIdVm,
	}
}

type ViewUpdatedIntEvent struct {
	Name   GameSocketIntEventName `json:"name"`
	GameId string                 `json:"gameId"`
}

func NewViewUpdatedIntEvent(gameIdVm string) ViewUpdatedIntEvent {
	return ViewUpdatedIntEvent{
		Name:   ViewUpdatedGameSocketIntEventName,
		GameId: gameIdVm,
	}
}
