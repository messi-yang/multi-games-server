package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType       ServerEventType = "ERRORED"
	GameJoinedServerEventType    ServerEventType = "GAME_JOINED"
	CameraChangedServerEventType ServerEventType = "CAMERA_CHANGED"
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
		Camera  viewmodel.CameraVm `json:"camera"`
		MapSize viewmodel.SizeVm   `json:"mapSize"`
		View    viewmodel.ViewVm   `json:"view"`
	} `json:"payload"`
}

type CameraChangedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Camera viewmodel.CameraVm `json:"camera"`
		View   viewmodel.ViewVm   `json:"view"`
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
