package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/identitymodel"
	"github.com/google/uuid"
)

type UserDto struct {
	Id           uuid.UUID `json:"id"`
	EmailAddress string    `json:"emailAddress"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewUserDto(user identitymodel.User) UserDto {
	dto := UserDto{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
		CreatedAt:    user.GetCreatedAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
	return dto
}
