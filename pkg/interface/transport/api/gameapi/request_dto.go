package gameapi

import "github.com/google/uuid"

type RequestDtoType string

const (
	pingRequestDtoType        RequestDtoType = "PING"
	moveRequestDtoType        RequestDtoType = "MOVE"
	placeItemRequestDtoType   RequestDtoType = "PLACE_ITEM"
	destroyItemRequestDtoType RequestDtoType = "DESTROY_ITEM"
)

type genericRequestDto struct {
	Type RequestDtoType `json:"type"`
}

type moveRequestDto struct {
	Type      RequestDtoType `json:"type"`
	Direction int8           `json:"direction"`
}

type placeItemRequestDto struct {
	Type   RequestDtoType `json:"type"`
	ItemId uuid.UUID      `json:"itemId"`
}

type destroyItemRequestDto struct {
	Type RequestDtoType `json:"type"`
}
