package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type PlayerModel struct {
	Id         uuid.UUID  `gorm:"primaryKey"`
	UserId     *uuid.UUID `gorm:""`
	User       *UserModel `gorm:"foreignKey:UserId;references:Id"`
	WorldId    uuid.UUID  `gorm:"not null"`
	Name       string     `gorm:"not null"`
	PosX       int        `gorm:"not null"`
	PosZ       int        `gorm:"not null"`
	Direction  int8       `gorm:"not null"`
	HeldItemId *uuid.UUID `gorm:""`
	CreatedAt  time.Time  `gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime;not null"`
}

func (PlayerModel) TableName() string {
	return "players"
}
