package gamesocketservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
)

type ResponseDtoType string

const (
	ErroredResponseDtoType        ResponseDtoType = "ERRORED"
	GameJoinedResponseDtoType     ResponseDtoType = "GAME_JOINED"
	PlayersUpdatedResponseDtoType ResponseDtoType = "PLAYERS_UPDATED"
	ViewUpdatedResponseDtoType    ResponseDtoType = "VIEW_UPDATED"
)

type ErroredResponseDto struct {
	Type          ResponseDtoType `json:"type"`
	ClientMessage string          `json:"clientMessage"`
}

type GameJoinedResponseDto struct {
	Type     ResponseDtoType      `json:"type"`
	PlayerId string               `json:"playerId"`
	Players  []viewmodel.PlayerVm `json:"players"`
	View     viewmodel.ViewVm     `json:"view"`
	Items    []viewmodel.ItemVm   `json:"items"`
}

type PlayersUpdatedResponseDto struct {
	Type    ResponseDtoType      `json:"type"`
	Players []viewmodel.PlayerVm `json:"players"`
}

type ViewUpdatedResponseDto struct {
	Type ResponseDtoType  `json:"type"`
	View viewmodel.ViewVm `json:"view"`
}
