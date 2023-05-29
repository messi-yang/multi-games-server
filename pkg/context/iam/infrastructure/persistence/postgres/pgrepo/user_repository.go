package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func newUserModel(user usermodel.User) pgmodel.UserModel {
	return pgmodel.UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
	}
}

func parseUserModel(userModel pgmodel.UserModel) usermodel.User {
	return usermodel.NewUser(sharedkernelmodel.NewUserId(userModel.Id), userModel.EmailAddress, userModel.Username)
}

type userRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewUserRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository usermodel.Repo) {
	return &userRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *userRepo) Add(user usermodel.User) error {
	userModel := newUserModel(user)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&userModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&user)
}

func (repo *userRepo) Get(userId sharedkernelmodel.UserId) (user usermodel.User, err error) {
	userModel := pgmodel.UserModel{Id: userId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&userModel).Error
	}); err != nil {
		return user, err
	}

	return parseUserModel(userModel), nil
}

func (repo *userRepo) FindUserByEmailAddress(emailAddress string) (user usermodel.User, userFound bool, err error) {
	userModels := []pgmodel.UserModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&userModels, pgmodel.UserModel{EmailAddress: emailAddress}).Error
	}); err != nil {
		return user, userFound, err
	}

	userFound = len(userModels) >= 1
	if !userFound {
		return user, false, nil
	}
	return parseUserModel(userModels[0]), true, nil
}
