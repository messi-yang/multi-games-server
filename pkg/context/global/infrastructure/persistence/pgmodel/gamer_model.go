package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type WorldAccountModel struct {
	Id               uuid.UUID `gorm:"primaryKey"`
	UserId           uuid.UUID `gorm:"unique;not null"`
	User             UserModel `gorm:"foreignKey:UserId;references:Id"`
	WorldsCount      int8      `gorm:"not null"`
	WorldsCountLimit int8      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;not null"`
}

func (WorldAccountModel) TableName() string {
	return "world_accounts"
}
