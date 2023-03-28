package gamesocketcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/google/uuid"
)

type responseDtoType string

const (
	playersUpdatedResponseDtoType responseDtoType = "PLAYERS_UPDATED"
	unitsUpdatedResponseDtoType   responseDtoType = "UNITS_UPDATED"
	gameJoinedResponseDtoType     responseDtoType = "GAME_JOINED"
)

type gameJoinedResponseDto struct {
	Type        responseDtoType    `json:"type"`
	PlayerId    uuid.UUID          `json:"playerId"`
	Players     []dto.PlayerAggDto `json:"players"`
	VisionBound dto.BoundVoDto     `json:"visionBound"`
	Units       []dto.UnitVoDto    `json:"units"`
	Items       []dto.ItemAggDto   `json:"items"`
}

type playersUpdatedResponseDto struct {
	Type    responseDtoType    `json:"type"`
	Players []dto.PlayerAggDto `json:"players"`
}

type unitsUpdatedResponseDto struct {
	Type        responseDtoType `json:"type"`
	VisionBound dto.BoundVoDto  `json:"visionBound"`
	Units       []dto.UnitVoDto `json:"units"`
}
