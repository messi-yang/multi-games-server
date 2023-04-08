package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/google/uuid"
)

type GamerAggDto struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
}

func NewGamerAggDto(gamer gamermodel.GamerAgg) GamerAggDto {
	return GamerAggDto{
		Id:     gamer.GetId().Uuid(),
		UserId: gamer.GetUserId().Uuid(),
	}
}
