package gamesocketservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType        ServerEventType = "ERRORED"
	GameJoinedServerEventType     ServerEventType = "GAME_JOINED"
	PlayersUpdatedServerEventType ServerEventType = "PLAYERS_UPDATED"
	ViewUpdatedServerEventType    ServerEventType = "VIEW_UPDATED"
)

type GenericServerEvent struct {
	Type ServerEventType `json:"type"`
}

type ErroredServerEvent struct {
	Type          ServerEventType `json:"type"`
	ClientMessage string          `json:"clientMessage"`
}

type GameJoinedServerEvent struct {
	Type     ServerEventType      `json:"type"`
	PlayerId string               `json:"playerId"`
	Players  []viewmodel.PlayerVm `json:"players"`
	View     viewmodel.ViewVm     `json:"view"`
	Items    []viewmodel.ItemVm   `json:"items"`
}

type PlayersUpdatedServerEvent struct {
	Type    ServerEventType      `json:"type"`
	Players []viewmodel.PlayerVm `json:"players"`
}

type ViewUpdatedServerEvent struct {
	Type ServerEventType  `json:"type"`
	View viewmodel.ViewVm `json:"view"`
}
