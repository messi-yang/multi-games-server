package gamesocketappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
)

type RequestDtoType string

const (
	PingRequestDtoType        RequestDtoType = "PING"
	MoveRequestDtoType        RequestDtoType = "MOVE"
	PlaceItemRequestDtoType   RequestDtoType = "PLACE_ITEM"
	DestroyItemRequestDtoType RequestDtoType = "DESTROY_ITEM"
)

type GenericRequestDto struct {
	Type RequestDtoType `json:"type"`
}

type MoveRequestDto struct {
	Type      RequestDtoType `json:"type"`
	Direction int8           `json:"direction"`
}

type PlaceItemRequestDto struct {
	Type       RequestDtoType  `json:"type"`
	Location   dto.LocationDto `json:"location"`
	ItemId     int16           `json:"itemId"`
	ActionedAt time.Time       `json:"actionedAt"`
}

type DestroyItemRequestDto struct {
	Type       RequestDtoType  `json:"type"`
	Location   dto.LocationDto `json:"location"`
	ActionedAt time.Time       `json:"actionedAt"`
}
