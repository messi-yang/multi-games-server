package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType       ServerEventType = "ERRORED"
	GameJoinedServerEventType    ServerEventType = "GAME_JOINED"
	CameraChangedServerEventType ServerEventType = "CAMERA_CHANGED"
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
		PlayerId string             `json:"playerId"`
		Camera   viewmodel.CameraVm `json:"camera"`
		Size     viewmodel.SizeVm   `json:"size"`
		View     viewmodel.ViewVm   `json:"view"`
	} `json:"payload"`
}

type CameraChangedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Camera viewmodel.CameraVm `json:"camera"`
		View   viewmodel.ViewVm   `json:"view"`
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
