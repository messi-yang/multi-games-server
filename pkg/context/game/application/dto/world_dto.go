package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldDto struct {
	Id      uuid.UUID `json:"id"`
	GamerId uuid.UUID `json:"gamerId"`
	Name    string    `json:"name"`
}

func NewWorldDto(world worldmodel.World) WorldDto {
	return WorldDto{
		Id:      world.GetId().Uuid(),
		GamerId: world.GetGamerId().Uuid(),
		Name:    world.GetName(),
	}
}
