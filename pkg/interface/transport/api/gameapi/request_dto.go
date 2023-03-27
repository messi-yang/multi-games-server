package gameapi

import "github.com/google/uuid"

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
	Type   RequestDtoType `json:"type"`
	ItemId uuid.UUID      `json:"itemId"`
}

type DestroyItemRequestDto struct {
	Type RequestDtoType `json:"type"`
}
