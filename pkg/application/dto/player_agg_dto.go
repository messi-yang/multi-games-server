package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"

type PlayerAggDto struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Position  PositionVoDto `json:"position"`
	Direction int8          `json:"direction"`
}

func NewPlayerAggDto(player playermodel.PlayerAgg) PlayerAggDto {
	return PlayerAggDto{
		Id:        player.GetId().String(),
		Name:      player.GetName(),
		Position:  NewPositionVoDto(player.GetPosition()),
		Direction: player.GetDirection().Int8(),
	}
}
