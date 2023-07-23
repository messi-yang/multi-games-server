package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/google/uuid"
)

type WorldMemberDto struct {
	Id        uuid.UUID   `json:"id"`
	User      dto.UserDto `json:"user"`
	WorldId   uuid.UUID   `json:"worldId"`
	Role      string      `json:"role"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func NewWorldMemberDto(worldMember worldaccessmodel.WorldMember, user usermodel.User) WorldMemberDto {
	dto := WorldMemberDto{
		Id:        worldMember.GetId().Uuid(),
		User:      dto.NewUserDto(user),
		WorldId:   worldMember.GeWorldId().Uuid(),
		Role:      worldMember.GetRole().String(),
		CreatedAt: worldMember.GetCreatedAt(),
		UpdatedAt: worldMember.GetUpdatedAt(),
	}
	return dto
}
