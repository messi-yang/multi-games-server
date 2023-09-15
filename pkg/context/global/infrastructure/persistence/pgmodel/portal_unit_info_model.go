package pgmodel

import (
	"github.com/google/uuid"
)

type PortalUnitInfoModel struct {
	Id         uuid.UUID `gorm:"not null"`
	WorldId    uuid.UUID `gorm:"not null"`
	TargetPosX *int      `gorm:""`
	TargetPosZ *int      `gorm:""`
}

func (PortalUnitInfoModel) TableName() string {
	return "portal_unit_infos"
}
