package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"

type PlayerDto struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Position  PositionDto `json:"position"`
	Direction int8        `json:"direction"`
}

func NewPlayerDto(player playermodel.PlayerAgg) PlayerDto {
	return PlayerDto{
		Id:        player.GetId().String(),
		Name:      player.GetName(),
		Position:  NewPositionDto(player.GetPosition()),
		Direction: player.GetDirection().ToInt8(),
	}
}
