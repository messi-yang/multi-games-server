package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gameaccountmodel"
	"github.com/google/uuid"
)

type GameAccountDto struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	RoomsCount int8      `json:"roomsCount"`
}

func NewGameAccountDto(gameAccount gameaccountmodel.GameAccount) GameAccountDto {
	return GameAccountDto{
		Id:         gameAccount.GetId().Uuid(),
		UserId:     gameAccount.GetUserId().Uuid(),
		RoomsCount: gameAccount.GetRoomsCount(),
	}
}
