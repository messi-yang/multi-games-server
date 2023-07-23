package pgmodel

import (
	"time"

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
