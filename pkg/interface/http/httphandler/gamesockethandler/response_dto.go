package gamesockethandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
)

type responseDtoType string

const (
	playersUpdatedResponseDtoType responseDtoType = "PLAYERS_UPDATED"
	unitsUpdatedResponseDtoType   responseDtoType = "UNITS_UPDATED"
	gameJoinedResponseDtoType     responseDtoType = "GAME_JOINED"
)

type gameJoinedResponseDto struct {
	Type responseDtoType `json:"type"`
}

type playersUpdatedResponseDto struct {
	Type         responseDtoType `json:"type"`
	MyPlayer     dto.PlayerDto   `json:"myPlayer"`
	OtherPlayers []dto.PlayerDto `json:"otherPlayers"`
}

type unitsUpdatedResponseDto struct {
	Type  responseDtoType `json:"type"`
	Units []dto.UnitDto   `json:"units"`
}
