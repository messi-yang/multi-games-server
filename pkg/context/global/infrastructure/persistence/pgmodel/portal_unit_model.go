package pgmodel

import (
	"github.com/google/uuid"
)

type PortalUnitModel struct {
	WorldId    uuid.UUID `gorm:"not null"`
	PosX       int       `gorm:"not null"`
	PosZ       int       `gorm:"not null"`
	ItemId     uuid.UUID `gorm:"not null"`
	Direction  int8      `gorm:"not null"`
	TargetPosX *int      `gorm:""`
	TargetPosZ *int      `gorm:""`
}

func (PortalUnitModel) TableName() string {
	return "portal_units"
}
