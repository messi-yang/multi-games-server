package pgmodel

import (
	"github.com/google/uuid"
)

type PortalUnitModel struct {
	Id         uuid.UUID `gorm:"not null"`
	TargetPosX *int      `gorm:""`
	TargetPosZ *int      `gorm:""`
}

func (PortalUnitModel) TableName() string {
	return "portal_units"
}
