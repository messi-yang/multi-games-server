package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
	"github.com/google/uuid"
)

type UserAggDto struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
}

func NewUserAggDto(user usermodel.UserAgg) UserAggDto {
	return UserAggDto{
		Id:     user.GetId().Uuid(),
		UserId: user.GetUserId().Uuid(),
	}
}
