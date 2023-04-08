package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldAggDto struct {
	Id      uuid.UUID `json:"id"`
	GamerId uuid.UUID `json:"gamerId"`
	Name    string    `json:"name"`
}

func NewWorldAggDto(world worldmodel.WorldAgg) WorldAggDto {
	return WorldAggDto{
		Id:      world.GetId().Uuid(),
		GamerId: world.GetGamerId().Uuid(),
		Name:    world.GetName(),
	}
}
