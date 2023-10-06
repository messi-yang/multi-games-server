package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/google/uuid"
)

type responseDtoType string

const (
	errorHappenedResponseType         responseDtoType = "ERROR_HAPPENED"
	worldEnteredResponseType          responseDtoType = "WORLD_ENTERED"
	staticUnitCreatedResponseType     responseDtoType = "STATIC_UNIT_CREATED"
	portalUnitCreatedResponseType     responseDtoType = "PORTAL_UNIT_CREATED"
	unitRotatedResponseType           responseDtoType = "UNIT_ROTATED"
	unitRemovedResponseType           responseDtoType = "UNIT_REMOVED"
	playerJoinedResponseType          responseDtoType = "PLAYER_JOINED"
	playerLeftResponseType            responseDtoType = "PLAYER_LEFT"
	playerMovedResponseType           responseDtoType = "PLAYER_MOVED"
	playerHeldItemChangedResponseType responseDtoType = "PLAYER_HELD_ITEM_CHANGED"
)

type errorHappenedResponse struct {
	Type    responseDtoType `json:"type"`
	Message string          `json:"message"`
}

type worldEnteredResponse struct {
	Type       responseDtoType          `json:"type"`
	World      viewmodel.WorldViewModel `json:"world"`
	Units      []dto.UnitDto            `json:"units"`
	MyPlayerId uuid.UUID                `json:"myPlayerId"`
	Players    []dto.PlayerDto          `json:"players"`
}

type staticUnitCreatedResponse struct {
	Type      responseDtoType `json:"type"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type portalUnitCreatedResponse struct {
	Type      responseDtoType `json:"type"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type unitRotatedResponse struct {
	Type     responseDtoType `json:"type"`
	Position dto.PositionDto `json:"position"`
}

type unitRemovedResponse struct {
	Type     responseDtoType `json:"type"`
	Position dto.PositionDto `json:"position"`
}

type playerJoinedResponse struct {
	Type   responseDtoType `json:"type"`
	Player dto.PlayerDto   `json:"player"`
}

type playerLeftResponse struct {
	Type     responseDtoType `json:"type"`
	PlayerId uuid.UUID       `json:"playerId"`
}

type playerMovedResponse struct {
	Type      responseDtoType `json:"type"`
	PlayerId  uuid.UUID       `json:"playerId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type playerHeldItemChangedResponse struct {
	Type     responseDtoType `json:"type"`
	PlayerId uuid.UUID       `json:"playerId"`
	ItemId   uuid.UUID       `json:"itemId"`
}
