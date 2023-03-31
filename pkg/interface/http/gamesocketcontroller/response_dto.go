package gamesocketcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httpdto"
	"github.com/google/uuid"
)

type responseDtoType string

const (
	playersUpdatedResponseDtoType responseDtoType = "PLAYERS_UPDATED"
	unitsUpdatedResponseDtoType   responseDtoType = "UNITS_UPDATED"
	gameJoinedResponseDtoType     responseDtoType = "GAME_JOINED"
)

type gameJoinedResponseDto struct {
	Type     responseDtoType      `json:"type"`
	PlayerId uuid.UUID            `json:"playerId"`
	Items    []httpdto.ItemAggDto `json:"items"`
}

type playersUpdatedResponseDto struct {
	Type    responseDtoType        `json:"type"`
	Players []httpdto.PlayerAggDto `json:"players"`
}

type unitsUpdatedResponseDto struct {
	Type        responseDtoType     `json:"type"`
	VisionBound httpdto.BoundVoDto  `json:"visionBound"`
	Units       []httpdto.UnitVoDto `json:"units"`
}
