package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/google/uuid"
)

type WorldAccountDto struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	WorldsCount int8      `json:"worldsCount"`
}

func NewWorldAccountDto(worldAccount worldaccountmodel.WorldAccount) WorldAccountDto {
	return WorldAccountDto{
		Id:          worldAccount.GetId().Uuid(),
		UserId:      worldAccount.GetUserId().Uuid(),
		WorldsCount: worldAccount.GetWorldsCount(),
	}
}
