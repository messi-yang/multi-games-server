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
	Type        ResponseDtoType `json:"type"`
	PlayerId    string          `json:"playerId"`
	Players     []dto.PlayerDto `json:"players"`
	VisionBound dto.BoundDto    `json:"visionBound"`
	Units       []dto.UnitDto   `json:"units"`
	Items       []dto.ItemDto   `json:"items"`
}

type PlayersUpdatedResponseDto struct {
	Type    ResponseDtoType `json:"type"`
	Players []dto.PlayerDto `json:"players"`
}

type UnitsUpdatedResponseDto struct {
	Type        ResponseDtoType `json:"type"`
	VisionBound dto.BoundDto    `json:"visionBound"`
	Units       []dto.UnitDto   `json:"units"`
}
