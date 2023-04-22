package gamesockethandler

import "github.com/google/uuid"

type RequestDtoType string

const (
	pingRequestDtoType           RequestDtoType = "PING"
	moveRequestDtoType           RequestDtoType = "MOVE"
	changeHeldItemRequestDtoType RequestDtoType = "CHANGE_HELD_ITEM"
	placeItemRequestDtoType      RequestDtoType = "PLACE_ITEM"
	removeItemRequestDtoType     RequestDtoType = "REMOVE_ITEM"
)

type genericRequestDto struct {
	Type RequestDtoType `json:"type"`
}

type moveRequestDto struct {
	Type      RequestDtoType `json:"type"`
	Direction int8           `json:"direction"`
}

type changeHeldItemRequestDto struct {
	Type   RequestDtoType `json:"type"`
	ItemId uuid.UUID      `json:"itemId"`
}

type placeItemRequestDto struct {
	Type RequestDtoType `json:"type"`
}

type removeItemRequestDto struct {
	Type RequestDtoType `json:"type"`
}
