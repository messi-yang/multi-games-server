package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamermodel"
	"github.com/google/uuid"
)

type GamerDto struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	WorldsCount int8      `json:"worldsCount"`
}

func NewGamerDto(gamer gamermodel.Gamer) GamerDto {
	return GamerDto{
		Id:          gamer.GetId().Uuid(),
		UserId:      gamer.GetUserId().Uuid(),
		WorldsCount: gamer.GetWorldsCount(),
	}
}
