package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type WorldModel struct {
	Id         uuid.UUID     `gorm:"primaryKey;unique"`
	GameUserId uuid.UUID     `gorm:"unique;not null"`
	GameUser   GameUserModel `gorm:"foreignKey:GameUserId;references:Id"`
	Name       string        `gorm:"not null"`
	CreatedAt  time.Time     `gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime;not null"`
}

func (WorldModel) TableName() string {
	return "worlds"
}
