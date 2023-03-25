package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"

type WorldAggDto struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func NewWorldAggDto(world worldmodel.WorldAgg) WorldAggDto {
	return WorldAggDto{
		Id:     world.GetId().String(),
		UserId: world.GetUserId().String(),
		Name:   world.GetName(),
	}
}
