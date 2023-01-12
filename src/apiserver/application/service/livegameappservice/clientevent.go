package livegameappservice

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ClientEventType string

const (
	NilClientEventType          ClientEventType = ""
	PingClientEventType         ClientEventType = "PING"
	ChangeCameraClientEventType ClientEventType = "CHANGE_CAMERA"
	BuildItemClientEventType    ClientEventType = "BUILD_ITEM"
	DestroyItemClientEventType  ClientEventType = "DESTROY_ITEM"
)

type GenericClientEvent struct {
	Type ClientEventType `json:"type"`
}

func ParseClientEventType(message []byte) (ClientEventType, error) {
	var clientEvent GenericClientEvent
	err := json.Unmarshal(message, &clientEvent)
	if err != nil {
		return NilClientEventType, err
	}

	return clientEvent.Type, nil
}

func ParseClientEvent[T any](message []byte) (T, error) {
	var clientEvent T
	err := json.Unmarshal(message, &clientEvent)
	if err != nil {
		return clientEvent, err
	}

	return clientEvent, nil
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
