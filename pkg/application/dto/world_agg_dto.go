package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldAggDto struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
}

func NewWorldAggDto(world worldmodel.WorldAgg) WorldAggDto {
	return WorldAggDto{
		Id:     world.GetId().Uuid(),
		UserId: world.GetUserId().Uuid(),
		Name:   world.GetName(),
	}
}
