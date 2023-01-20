package livegameappservice

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
