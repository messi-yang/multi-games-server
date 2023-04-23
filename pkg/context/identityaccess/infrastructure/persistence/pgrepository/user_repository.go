package pgrepository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"gorm.io/gorm"
)

func newUserModel(user usermodel.UserAgg) pgmodel.UserModel {
	return pgmodel.UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
	}
}

func parseUserModel(userModel pgmodel.UserModel) usermodel.UserAgg {
	return usermodel.NewUserAgg(sharedkernelmodel.NewUserIdVo(userModel.Id), userModel.EmailAddress, userModel.Username)
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

func (repo *userRepository) Add(user usermodel.UserAgg) error {
	userModel := newUserModel(user)
	res := repo.dbClient.Create(&userModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *userRepository) Get(userId sharedkernelmodel.UserIdVo) (user usermodel.UserAgg, err error) {
	userModel := pgmodel.UserModel{Id: userId.Uuid()}
	result := repo.dbClient.First(&userModel)
	if result.Error != nil {
		return user, result.Error
	}

	return parseUserModel(userModel), nil
}

func (repo *userRepository) FindUserByEmailAddress(emailAddress string) (user usermodel.UserAgg, userFound bool, err error) {
	userModels := []pgmodel.UserModel{}
	result := repo.dbClient.Find(&userModels, pgmodel.UserModel{EmailAddress: emailAddress})
	if result.Error != nil {
		return user, userFound, result.Error
	}
	userFound = result.RowsAffected >= 1
	if !userFound {
		return user, false, nil
	}
	return parseUserModel(userModels[0]), true, nil
}
