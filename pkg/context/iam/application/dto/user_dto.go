package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/google/uuid"
)

type UserDto struct {
	Id           uuid.UUID `json:"id"`
	EmailAddress string    `json:"emailAddress"`
	Username     string    `json:"username"`
}

func NewUserDto(user usermodel.User) UserDto {
	dto := UserDto{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
	}
	return dto
}
