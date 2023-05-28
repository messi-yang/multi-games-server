package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type WorldRoleName string

const (
	WorldRoleAdmin WorldRoleName = "admin"
)

type WorldRoleModel struct {
	Id        uuid.UUID     `gorm:"primaryKey"`
	UserId    uuid.UUID     `gorm:"not null"`
	WorldId   uuid.UUID     `gorm:"not null"`
	Name      WorldRoleName `gorm:"not null"`
	CreatedAt time.Time     `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime;not null"`
}

func (WorldRoleModel) TableName() string {
	return "world_roles"
}
