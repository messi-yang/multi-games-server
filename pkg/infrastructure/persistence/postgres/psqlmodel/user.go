package psqlmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
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

func NewUserModel(user usermodel.UserAgg) UserModel {
	return UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
	}
}

func (model UserModel) ToAggregate() usermodel.UserAgg {
	return usermodel.NewUserAgg(usermodel.NewUserIdVo(model.Id), model.EmailAddress, model.Username)
}
