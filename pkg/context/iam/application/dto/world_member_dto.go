package dto

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/google/uuid"
)

type WorldMemberDto struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	WorldId   uuid.UUID `json:"worldId"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewWorldMemberDto(worldMember worldaccessmodel.WorldMember) WorldMemberDto {
	dto := WorldMemberDto{
		Id:        worldMember.GetId().Uuid(),
		UserId:    worldMember.GeUserId().Uuid(),
		WorldId:   worldMember.GeWorldId().Uuid(),
		Role:      worldMember.GetRole().String(),
		CreatedAt: worldMember.GetCreatedAt(),
		UpdatedAt: worldMember.GetUpdatedAt(),
	}
	return dto
}
