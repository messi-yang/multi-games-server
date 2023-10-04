package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type RequestType string

const (
	pingRequestType                 RequestType = "PING"
	moveRequestType                 RequestType = "MOVE"
	changePlayerHeldItemRequestType RequestType = "CHANGE_PLAYER_HELD_ITEM"
	createStaticUnitRequestType     RequestType = "CREATE_STATIC_UNIT"
	createPortalUnitRequestType     RequestType = "CREATE_PORTAL_UNIT"
	rotateUnitRequestType           RequestType = "ROTATE_UNIT"
	removeUnitRequestType           RequestType = "REMOVE_UNIT"
)

type genericRequest struct {
	Type RequestType `json:"type"`
}

type moveRequest struct {
	Type      RequestType `json:"type"`
	Direction int8        `json:"direction"`
}

type changePlayerHeldItemRequest struct {
	Type     RequestType `json:"type"`
	PlayerId uuid.UUID   `json:"playerId"`
	ItemId   uuid.UUID   `json:"itemId"`
}

type createStaticUnitRequest struct {
	Type      RequestType     `json:"type"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type createPortalUnitRequest struct {
	Type      RequestType     `json:"type"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type rotateUnitRequest struct {
	Type     RequestType     `json:"type"`
	Position dto.PositionDto `json:"position"`
}

type removeUnitRequest struct {
	Type     RequestType     `json:"type"`
	Position dto.PositionDto `json:"position"`
}
