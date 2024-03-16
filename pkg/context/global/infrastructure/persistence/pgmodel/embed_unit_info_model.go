package pgmodel

import (
	"github.com/google/uuid"
)

type EmbedUnitInfoModel struct {
	Id        uuid.UUID `gorm:"not null"`
	WorldId   uuid.UUID `gorm:"not null"`
	EmbedCode string    `gorm:"not null"`
}

func (EmbedUnitInfoModel) TableName() string {
	return "embed_unit_infos"
}
