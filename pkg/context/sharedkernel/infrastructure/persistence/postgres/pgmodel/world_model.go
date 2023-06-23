package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type WorldModel struct {
	Id         uuid.UUID `gorm:"primaryKey"`
	UserId     uuid.UUID `gorm:"unique;not null"`
	Name       string    `gorm:"not null"`
	BoundFromX int       `gorm:"not null"`
	BoundFromZ int       `gorm:"not null"`
	BoundToX   int       `gorm:"not null"`
	BoundToZ   int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;not null"`
}

func (WorldModel) TableName() string {
	return "worlds"
}
