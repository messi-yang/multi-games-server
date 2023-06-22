package worldjourneyhandler

import "github.com/google/uuid"

type RequestType string

const (
	pingRequestType           RequestType = "PING"
	moveRequestType           RequestType = "MOVE"
	changeHeldItemRequestType RequestType = "CHANGE_HELD_ITEM"
	placeItemRequestType      RequestType = "PLACE_ITEM"
	removeItemRequestType     RequestType = "REMOVE_ITEM"
)

type genericRequest struct {
	Type RequestType `json:"type"`
}

type moveRequest struct {
	Type      RequestType `json:"type"`
	Direction int8        `json:"direction"`
}

type changeHeldItemRequest struct {
	Type   RequestType `json:"type"`
	ItemId uuid.UUID   `json:"itemId"`
}

type placeItemRequest struct {
	Type RequestType `json:"type"`
}

type removeItemRequest struct {
	Type RequestType `json:"type"`
}
