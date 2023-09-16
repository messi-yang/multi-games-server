package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/google/uuid"
)

type UserDto struct {
	Id           uuid.UUID `json:"id"`
	EmailAddress string    `json:"emailAddress"`
	Username     string    `json:"username"`
	FriendlyName string    `json:"friendlyName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewUserDto(user usermodel.User) UserDto {
	dto := UserDto{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress().String(),
		Username:     user.GetUsername().String(),
		FriendlyName: user.GetFriendlyName().String(),
		CreatedAt:    user.GetCreatedAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
	return dto
}
