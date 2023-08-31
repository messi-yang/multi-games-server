package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type RequestType string

const (
	pingRequestType             RequestType = "PING"
	moveRequestType             RequestType = "MOVE"
	changeHeldItemRequestType   RequestType = "CHANGE_HELD_ITEM"
	createStaticUnitRequestType RequestType = "CREATE_STATIC_UNIT"
	removeUnitRequestType       RequestType = "REMOVE_UNIT"
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

type createStaticUnitRequest struct {
	Type      RequestType     `json:"type"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removeUnitRequest struct {
	Type     RequestType     `json:"type"`
	Position dto.PositionDto `json:"position"`
}
