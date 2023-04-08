package pgmodel

import (
	"time"

	"github.com/google/uuid"
)

type GameUserModel struct {
	Id        uuid.UUID `gorm:"primaryKey;unique"`
	UserId    uuid.UUID `gorm:"unique;not null"`
	User      UserModel `gorm:"foreignKey:UserId;references:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

func (GameUserModel) TableName() string {
	return "game_users"
}
