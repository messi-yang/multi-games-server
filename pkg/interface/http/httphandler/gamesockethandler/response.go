package gamesockethandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
)

type responseDtoType string

const (
	errorHappenedResponseType  responseDtoType = "ERROR_HAPPENED"
	worldEnteredResponseType   responseDtoType = "WORLD_ENTERED"
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

type playersUpdatedResponse struct {
	Type         responseDtoType `json:"type"`
	MyPlayer     dto.PlayerDto   `json:"myPlayer"`
	OtherPlayers []dto.PlayerDto `json:"otherPlayers"`
}

type unitsUpdatedResponse struct {
	Type  responseDtoType `json:"type"`
	Units []dto.UnitDto   `json:"units"`
}
