package dto

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/accessmodel"
	"github.com/google/uuid"
)

type UserWorldRoleDto struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	WorldId   uuid.UUID `json:"worldId"`
	WorldRole string    `json:"worldRole"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserWorldRoleDto(userWorldRole accessmodel.UserWorldRole) UserWorldRoleDto {
	dto := UserWorldRoleDto{
		Id:        userWorldRole.GetId().Uuid(),
		UserId:    userWorldRole.GeUserId().Uuid(),
		WorldId:   userWorldRole.GeWorldId().Uuid(),
		WorldRole: userWorldRole.GetWorldRole().String(),
		CreatedAt: userWorldRole.GetCreatedAt(),
		UpdatedAt: userWorldRole.GetUpdatedAt(),
	}
	return dto
}
