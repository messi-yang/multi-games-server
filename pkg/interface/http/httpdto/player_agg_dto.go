package httpdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/google/uuid"
)

type PlayerAggDto struct {
	Id        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	Position  PositionVoDto `json:"position"`
	Direction int8          `json:"direction"`
}

func NewPlayerAggDto(player playermodel.PlayerAgg) PlayerAggDto {
	return PlayerAggDto{
		Id:        player.GetId().Uuid(),
		Name:      player.GetName(),
		Position:  NewPositionVoDto(player.GetPosition()),
		Direction: player.GetDirection().Int8(),
	}
}
