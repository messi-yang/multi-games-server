package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type GamerModel struct {
	Id               uuid.UUID `gorm:"primaryKey"`
	UserId           uuid.UUID `gorm:"unique;not null"`
	User             UserModel `gorm:"foreignKey:UserId;references:Id"`
	WorldsCount      int8      `gorm:"not null"`
	WorldsCountLimit int8      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;not null"`
}

func (GamerModel) TableName() string {
	return "gamers"
}
