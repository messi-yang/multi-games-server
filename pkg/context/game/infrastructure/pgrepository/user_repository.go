package pgrepository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/infrastructure/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
	identityaccessusermodel "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newGameUserModel(user usermodel.UserAgg) pgmodel.GameUserModel {
	return pgmodel.GameUserModel{
		Id:     user.GetId().Uuid(),
		UserId: user.GetUserId().Uuid(),
	}
}

func parseGameUserModel(userModel pgmodel.GameUserModel) usermodel.UserAgg {
	return usermodel.NewUserAgg(
		usermodel.NewUserIdVo(userModel.Id),
		identityaccessusermodel.NewUserIdVo(userModel.UserId),
	)
}

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

func (repo *userRepository) GetAll() (users []usermodel.UserAgg, err error) {
	var userModels []pgmodel.GameUserModel
	result := repo.dbClient.Find(&userModels)
	if result.Error != nil {
		err = result.Error
		return users, err
	}

	users = lo.Map(userModels, func(userModel pgmodel.GameUserModel, _ int) usermodel.UserAgg {
		return parseGameUserModel(userModel)
	})
	return users, nil
}
