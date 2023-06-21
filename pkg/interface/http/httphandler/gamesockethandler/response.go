package gamesockethandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
)

type responseDtoType string

const (
	errorHappenedResponseType  responseDtoType = "ERROR_HAPPENED"
	worldEnteredResponseType   responseDtoType = "WORLD_ENTERED"
	unitCreatedResponseType    responseDtoType = "UNIT_CREATED"
	unitDeletedResponseType    responseDtoType = "UNIT_DELETED"
	playersUpdatedResponseType responseDtoType = "PLAYERS_UPDATED"
	unitsUpdatedResponseType   responseDtoType = "UNITS_UPDATED"
)

type errorHappenedResponse struct {
	Type    responseDtoType `json:"type"`
	Message string          `json:"message"`
}

type worldsEnteredResponse struct {
	Type         responseDtoType `json:"type"`
	World        dto.WorldDto    `json:"world"`
	Units        []dto.UnitDto   `json:"units"`
	MyPlayer     dto.PlayerDto   `json:"myPlayer"`
	OtherPlayers []dto.PlayerDto `json:"otherPlayers"`
}

type unitCreatedResponse struct {
	Type responseDtoType `json:"type"`
	Unit dto.UnitDto     `json:"unit"`
}

type unitDeletedResponse struct {
	Type     responseDtoType `json:"type"`
	Position dto.PositionDto `json:"position"`
}

type playersUpdatedResponse struct {
	Type         responseDtoType `json:"type"`
	MyPlayer     dto.PlayerDto   `json:"myPlayer"`
	OtherPlayers []dto.PlayerDto `json:"otherPlayers"`
}

type unitsUpdatedResponse struct {
	Type  responseDtoType `json:"type"`
	Units []dto.UnitDto   `json:"units"`
}
