package appservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ClientEventType string

const (
	PingClientEventType        ClientEventType = "PING"
	MoveClientEventType        ClientEventType = "MOVE"
	BuildItemClientEventType   ClientEventType = "BUILD_ITEM"
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

type BuildItemClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Location   viewmodel.LocationVm `json:"location"`
		ItemId     string               `json:"itemId"`
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
	ItemsUpdatedServerEventType   ServerEventType = "ITEMS_UPDATED"
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
		MyPlayer     viewmodel.PlayerVm   `json:"myPlayer"`
		OtherPlayers []viewmodel.PlayerVm `json:"otherPlayers"`
		MapSize      viewmodel.SizeVm     `json:"mapSize"`
		View         viewmodel.ViewVm     `json:"view"`
	} `json:"payload"`
}

type PlayersUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		MyPlayer     viewmodel.PlayerVm   `json:"myPlayer"`
		OtherPlayers []viewmodel.PlayerVm `json:"otherPlayers"`
	} `json:"payload"`
}

type ViewUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		View viewmodel.ViewVm `json:"view"`
	} `json:"payload"`
}

type ItemsUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Items []viewmodel.ItemVm `json:"items"`
	} `json:"payload"`
}
