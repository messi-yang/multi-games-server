package gamesocketcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/jsondto"
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
	Type         responseDtoType        `json:"type"`
	MyPlayer     jsondto.PlayerAggDto   `json:"myPlayer"`
	OtherPlayers []jsondto.PlayerAggDto `json:"otherPlayers"`
}

type unitsUpdatedResponseDto struct {
	Type        responseDtoType     `json:"type"`
	VisionBound jsondto.BoundVoDto  `json:"visionBound"`
	Units       []jsondto.UnitVoDto `json:"units"`
}
