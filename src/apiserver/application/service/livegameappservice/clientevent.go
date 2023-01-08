package livegameappservice

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ClientEventType string

const (
	NilClientEventType           ClientEventType = ""
	PingClientEventType          ClientEventType = "PING"
	ObserveExtentClientEventType ClientEventType = "OBSERVE_MAP_RANGE"
	BuildItemClientEventType     ClientEventType = "BUILD_ITEM"
	DestroyItemClientEventType   ClientEventType = "DESTROY_ITEM"
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

type BuildItemClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Location   viewmodel.LocationViewModel `json:"location"`
		ItemId     string                      `json:"itemId"`
		ActionedAt time.Time                   `json:"actionedAt"`
	} `json:"payload"`
}

type DestroyItemClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Location   viewmodel.LocationViewModel `json:"location"`
		ActionedAt time.Time                   `json:"actionedAt"`
	} `json:"payload"`
}

type ObserveExtentClientEvent struct {
	Type    ClientEventType `json:"type"`
	Payload struct {
		Extent     viewmodel.ExtentViewModel `json:"extent"`
		ActionedAt time.Time                 `json:"actionedAt"`
	} `json:"payload"`
}
