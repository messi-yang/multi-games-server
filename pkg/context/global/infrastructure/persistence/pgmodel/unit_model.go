package pgmodel

import (
	"github.com/google/uuid"
)

type UnitModel struct {
	WorldId      uuid.UUID    `gorm:"not null"`
	PosX         int          `gorm:"not null"`
	PosZ         int          `gorm:"not null"`
	ItemId       uuid.UUID    `gorm:"not null"`
	Direction    int8         `gorm:"not null"`
	Type         UnitTypeEnum `gorm:"not null"`
	LinkedUnitId *uuid.UUID   `gorm:"not null"`
}

func (UnitModel) TableName() string {
	return "units"
}
