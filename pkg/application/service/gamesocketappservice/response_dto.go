package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
)

type ResponseDtoType string

const (
	ErroredResponseDtoType        ResponseDtoType = "ERRORED"
	GameJoinedResponseDtoType     ResponseDtoType = "GAME_JOINED"
	PlayersUpdatedResponseDtoType ResponseDtoType = "PLAYERS_UPDATED"
	UnitsUpdatedResponseDtoType   ResponseDtoType = "UNITS_UPDATED"
)

type ErroredResponseDto struct {
	Type          ResponseDtoType `json:"type"`
	ClientMessage string          `json:"clientMessage"`
}

type GameJoinedResponseDto struct {
	Type        ResponseDtoType    `json:"type"`
	PlayerId    string             `json:"playerId"`
	Players     []dto.PlayerAggDto `json:"players"`
	VisionBound dto.BoundVoDto     `json:"visionBound"`
	Units       []dto.UnitVoDto    `json:"units"`
	Items       []dto.ItemAggDto   `json:"items"`
}

type PlayersUpdatedResponseDto struct {
	Type    ResponseDtoType    `json:"type"`
	Players []dto.PlayerAggDto `json:"players"`
}

type UnitsUpdatedResponseDto struct {
	Type        ResponseDtoType `json:"type"`
	VisionBound dto.BoundVoDto  `json:"visionBound"`
	Units       []dto.UnitVoDto `json:"units"`
}
