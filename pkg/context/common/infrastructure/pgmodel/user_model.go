package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id           uuid.UUID `gorm:"primaryKey;unique"`
	EmailAddress string    `gorm:"unique;not null"`
	Username     string    `gorm:"unique;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;not null"`
}

func (UserModel) TableName() string {
	return "users"
}
