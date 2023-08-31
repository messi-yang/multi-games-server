package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type ItemModel struct {
	Id                 uuid.UUID    `gorm:"primaryKey"`
	CompatibleUnitType UnitTypeEnum `gorm:"not null"`
	Name               string       `gorm:"not null"`
	Traversable        bool         `gorm:"not null"`
	ModelSrc           string       `gorm:"not null"`
	ThumbnailSrc       string       `gorm:"not null"`
	CreatedAt          time.Time    `gorm:"autoCreateTime;not null"`
	UpdatedAt          time.Time    `gorm:"autoUpdateTime;not null"`
}

func (ItemModel) TableName() string {
	return "items"
}
