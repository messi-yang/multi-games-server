package appservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ClientEventType string

const (
	PingClientEventType         ClientEventType = "PING"
	ChangeCameraClientEventType ClientEventType = "CHANGE_CAMERA"
	BuildItemClientEventType    ClientEventType = "BUILD_ITEM"
	DestroyItemClientEventType  ClientEventType = "DESTROY_ITEM"
)

type GenericClientEvent struct {
	Type ClientEventType `json:"type"`
}

type PingClientEvent struct {
	Type ClientEventType `json:"type"`
}

type ChangeCameraClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Camera viewmodel.CameraVm `json:"camera"`
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
	ErroredServerEventType       ServerEventType = "ERRORED"
	GameJoinedServerEventType    ServerEventType = "GAME_JOINED"
	PlayerUpdatedServerEventType ServerEventType = "PLAYER_UPDATED"
	ViewChangedServerEventType   ServerEventType = "VIEW_CHANGED"
	ViewUpdatedServerEventType   ServerEventType = "VIEW_UPDATED"
	ItemsUpdatedServerEventType  ServerEventType = "ITEMS_UPDATED"
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
		Player  viewmodel.PlayerVm `json:"player"`
		MapSize viewmodel.SizeVm   `json:"mapSize"`
		View    viewmodel.ViewVm   `json:"view"`
	} `json:"payload"`
}

type PlayerUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Player viewmodel.PlayerVm `json:"player"`
	} `json:"payload"`
}

type ViewChangedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		View viewmodel.ViewVm `json:"view"`
	} `json:"payload"`
}

type ViewUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Camera viewmodel.CameraVm `json:"camera"`
		View   viewmodel.ViewVm   `json:"view"`
	} `json:"payload"`
}

type ItemsUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Items []viewmodel.ItemVm `json:"items"`
	} `json:"payload"`
}
