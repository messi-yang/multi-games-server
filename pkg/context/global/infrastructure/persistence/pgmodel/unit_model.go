package pgmodel

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type UnitModel struct {
	WorldId      uuid.UUID     `gorm:"not null"`
	PosX         int           `gorm:"not null"`
	PosZ         int           `gorm:"not null"`
	ItemId       uuid.UUID     `gorm:"not null"`
	Direction    int8          `gorm:"not null"`
	Type         UnitTypeEnum  `gorm:"not null"`
	InfoId       *uuid.UUID    `gorm:"not null"`
	InfoSnapshot *pgtype.JSONB `gorm:"type:jsonb"`
}

func (UnitModel) TableName() string {
	return "units"
}
