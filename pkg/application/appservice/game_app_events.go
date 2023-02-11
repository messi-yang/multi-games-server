package appservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
)

type ClientEventType string

const (
	PingClientEventType        ClientEventType = "PING"
	MoveClientEventType        ClientEventType = "MOVE"
	PlaceItemClientEventType   ClientEventType = "PLACE_ITEM"
	DestroyItemClientEventType ClientEventType = "DESTROY_ITEM"
)

type GenericClientEvent struct {
	Type ClientEventType `json:"type"`
}

type PingClientEvent struct {
	Type ClientEventType `json:"type"`
}

type MoveClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Direction int8 `json:"direction"`
	} `json:"payload"`
}

type PlaceItemClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Location   viewmodel.LocationVm `json:"location"`
		ItemId     int16                `json:"itemId"`
		ActionedAt time.Time            `json:"actionedAt"`
	} `json:"payload"`
}

type DestroyItemClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Location   viewmodel.LocationVm `json:"location"`
		ActionedAt time.Time            `json:"actionedAt"`
	} `json:"payload"`
}

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
		MapSize  viewmodel.SizeVm     `json:"mapSize"`
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
