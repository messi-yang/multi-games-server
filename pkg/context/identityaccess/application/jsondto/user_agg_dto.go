package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/google/uuid"
)

type UserAggDto struct {
	Id           uuid.UUID `json:"id"`
	EmailAddress string    `json:"emailAddress"`
	Username     string    `json:"username"`
}

func NewUserAggDto(user usermodel.UserAgg) UserAggDto {
	dto := UserAggDto{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
	}
	return dto
}
