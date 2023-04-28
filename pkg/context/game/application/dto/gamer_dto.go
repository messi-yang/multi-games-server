package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/google/uuid"
)

type GamerDto struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
}

func NewGamerDto(gamer gamermodel.Gamer) GamerDto {
	return GamerDto{
		Id:     gamer.GetId().Uuid(),
		UserId: gamer.GetUserId().Uuid(),
	}
}
