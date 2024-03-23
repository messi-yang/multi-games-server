package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/google/uuid"
)

type WorldRole string

const (
	WorldRoleOwner  WorldRole = "owner"
	WorldRoleAdmin  WorldRole = "admin"
	WorldRoleEditor WorldRole = "editor"
	WorldRoleViewer WorldRole = "viewer"
)

type WorldMemberModel struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	UserId    uuid.UUID `gorm:"not null"`
	WorldId   uuid.UUID `gorm:"not null"`
	Role      WorldRole `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

func (WorldMemberModel) TableName() string {
	return "world_members"
}

func NewWorldMemberModel(worldMember worldaccessmodel.WorldMember) WorldMemberModel {
	return WorldMemberModel{
		Id:        worldMember.GetId().Uuid(),
		WorldId:   worldMember.GeWorldId().Uuid(),
		UserId:    worldMember.GeUserId().Uuid(),
		Role:      WorldRole(worldMember.GetRole().String()),
		CreatedAt: worldMember.GetCreatedAt(),
		UpdatedAt: worldMember.GetUpdatedAt(),
	}
}

func ParseWorldMemberModel(worldMemberModel WorldMemberModel) (worldMember worldaccessmodel.WorldMember, err error) {
	worldRole, err := globalcommonmodel.NewWorldRole(string(worldMemberModel.Role))
	if err != nil {
		return worldMember, err
	}
	return worldaccessmodel.LoadWorldMember(
		worldaccessmodel.NewWorldMemberId(worldMemberModel.Id),
		globalcommonmodel.NewWorldId(worldMemberModel.WorldId),
		globalcommonmodel.NewUserId(worldMemberModel.UserId),
		worldRole,
		worldMemberModel.CreatedAt,
		worldMemberModel.UpdatedAt,
	), nil
}
