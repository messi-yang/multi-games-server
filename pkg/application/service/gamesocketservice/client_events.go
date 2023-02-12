package gamesocketservice

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
