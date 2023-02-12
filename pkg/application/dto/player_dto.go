package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"

type PlayerDto struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Location LocationDto `json:"location"`
}

func NewPlayerDto(player gamemodel.PlayerEntity) PlayerDto {
	return PlayerDto{
		Id:       player.GetId().ToString(),
		Name:     player.GetName(),
		Location: NewLocationDto(player.GetLocation()),
	}
}
