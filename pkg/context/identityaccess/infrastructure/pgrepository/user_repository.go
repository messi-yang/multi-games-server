package pgrepository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/infrastructure/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"gorm.io/gorm"
)

func newUserModel(user usermodel.UserAgg) pgmodel.UserModel {
	return pgmodel.UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
	}
}

// func parseUserModel(userModel pgmodel.UserModel) usermodel.UserAgg {
// 	return usermodel.NewUserAgg(usermodel.NewUserIdVo(userModel.Id), userModel.EmailAddress, userModel.Username)
// }

type userRepository struct {
	dbClient *gorm.DB
}

func NewUserRepository() (repository usermodel.Repository, err error) {
	dbClient, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &userRepository{dbClient: dbClient}, nil
}

func (repo *userRepository) Add(user usermodel.UserAgg) error {
	userModel := newUserModel(user)
	res := repo.dbClient.Create(&userModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
