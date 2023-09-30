package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/google/uuid"
)

type responseDtoType string

const (
	errorHappenedResponseType responseDtoType = "ERROR_HAPPENED"
	worldEnteredResponseType  responseDtoType = "WORLD_ENTERED"
	unitCreatedResponseType   responseDtoType = "UNIT_CREATED"
	unitUpdatedResponseType   responseDtoType = "UNIT_UPDATED"
	unitRemovedResponseType   responseDtoType = "UNIT_REMOVED"
	playerJoinedResponseType  responseDtoType = "PLAYER_JOINED"
	playerLeftResponseType    responseDtoType = "PLAYER_LEFT"
	playerMovedResponseType   responseDtoType = "PLAYER_MOVED"
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

type unitCreatedResponse struct {
	Type responseDtoType `json:"type"`
	Unit dto.UnitDto     `json:"unit"`
}

type unitUpdatedResponse struct {
	Type responseDtoType `json:"type"`
	Unit dto.UnitDto     `json:"unit"`
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
	Type   responseDtoType `json:"type"`
	Player dto.PlayerDto   `json:"player"`
}
