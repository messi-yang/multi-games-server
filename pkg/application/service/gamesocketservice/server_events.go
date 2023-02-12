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
	Type    ServerEventType `json:"type"`
	Payload struct {
		ClientMessage string `json:"clientMessage"`
	} `json:"payload"`
}

type GameJoinedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		PlayerId string               `json:"playerId"`
		Players  []viewmodel.PlayerVm `json:"players"`
		View     viewmodel.ViewVm     `json:"view"`
		Items    []viewmodel.ItemVm   `json:"items"`
	} `json:"payload"`
}

type PlayersUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Players []viewmodel.PlayerVm `json:"players"`
	} `json:"payload"`
}

type ViewUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		View viewmodel.ViewVm `json:"view"`
	} `json:"payload"`
}
