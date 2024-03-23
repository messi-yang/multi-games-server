package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/google/uuid"
)

type UserModel struct {
	Id           uuid.UUID `gorm:"primaryKey"`
	EmailAddress string    `gorm:"unique;not null"`
	Username     string    `gorm:"unique;not null"`
	FriendlyName string    `gorm:"unique;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;not null"`
}

func (UserModel) TableName() string {
	return "users"
}

func NewUserModel(user usermodel.User) UserModel {
	return UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress().String(),
		Username:     user.GetUsername().String(),
		FriendlyName: user.GetFriendlyName().String(),
		CreatedAt:    user.GetCreatedAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
}

func ParseUserModel(userModel UserModel) (user usermodel.User, err error) {
	emailAddress, err := globalcommonmodel.NewEmailAddress(userModel.EmailAddress)
	if err != nil {
		return user, err
	}
	username, err := globalcommonmodel.NewUsername(userModel.Username)
	if err != nil {
		return user, err
	}
	friendlyName, err := usermodel.NewFriendlyName(userModel.FriendlyName)
	if err != nil {
		return user, err
	}
	return usermodel.LoadUser(
		globalcommonmodel.NewUserId(userModel.Id),
		emailAddress,
		username,
		friendlyName,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	), nil
}
