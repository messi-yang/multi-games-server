package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/identitymodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func newUserModel(user identitymodel.User) pgmodel.UserModel {
	return pgmodel.UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress(),
		Username:     user.GetUsername(),
		CreatedAt:    user.GetCreatedAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
}

func parseUserModel(userModel pgmodel.UserModel) identitymodel.User {
	return identitymodel.LoadUser(
		sharedkernelmodel.NewUserId(userModel.Id),
		userModel.EmailAddress,
		userModel.Username,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	)
}

type userRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewUserRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository identitymodel.UserRepo) {
	return &userRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *userRepo) Add(user identitymodel.User) error {
	userModel := newUserModel(user)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&userModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&user)
}

func (repo *userRepo) Get(userId sharedkernelmodel.UserId) (user identitymodel.User, err error) {
	userModel := pgmodel.UserModel{Id: userId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&userModel).Error
	}); err != nil {
		return user, err
	}

	return parseUserModel(userModel), nil
}

func (repo *userRepo) FindUserByEmailAddress(emailAddress string) (user identitymodel.User, userFound bool, err error) {
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
