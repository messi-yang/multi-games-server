package worldjourneyhandler

import "github.com/google/uuid"

type RequestType string

const (
	pingRequestType           RequestType = "PING"
	moveRequestType           RequestType = "MOVE"
	changeHeldItemRequestType RequestType = "CHANGE_HELD_ITEM"
	placeUnitRequestType      RequestType = "PLACE_UNIT"
	removeUnitRequestType     RequestType = "REMOVE_UNIT"
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

type placeUnitRequest struct {
	Type RequestType `json:"type"`
}

type removeUnitRequest struct {
	Type RequestType `json:"type"`
}
