package gamesockethandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
)

type responseDtoType string

const (
	errorHappenedResponseType  responseDtoType = "ERROR_HAPPENED"
	playersUpdatedResponseType responseDtoType = "PLAYERS_UPDATED"
	unitsUpdatedResponseType   responseDtoType = "UNITS_UPDATED"
)

type errorHappenedResponse struct {
	Type    responseDtoType `json:"type"`
	Message string          `json:"message"`
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
