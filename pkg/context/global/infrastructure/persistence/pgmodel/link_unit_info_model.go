package pgmodel

import (
	"github.com/google/uuid"
)

type LinkUnitInfoModel struct {
	Id      uuid.UUID `gorm:"not null"`
	WorldId uuid.UUID `gorm:"not null"`
	Url     string    `gorm:"not null"`
}

func (LinkUnitInfoModel) TableName() string {
	return "link_unit_infos"
}
