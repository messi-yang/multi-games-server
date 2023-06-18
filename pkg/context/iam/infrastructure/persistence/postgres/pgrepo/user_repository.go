package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/identitymodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func newUserModel(user identitymodel.User) pgmodel.UserModel {
	return pgmodel.UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress().String(),
		Username:     user.GetUsername().String(),
		CreatedAt:    user.GetCreatedAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
}

func parseUserModel(userModel pgmodel.UserModel) (user identitymodel.User, err error) {
	emailAddress, err := sharedkernelmodel.NewEmailAddress(userModel.EmailAddress)
	if err != nil {
		return user, err
	}
	username, err := sharedkernelmodel.NewUsername(userModel.Username)
	if err != nil {
		return user, err
	}
	return identitymodel.LoadUser(
		sharedkernelmodel.NewUserId(userModel.Id),
		emailAddress,
		username,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	), nil
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

	return parseUserModel(userModel)
}

func (repo *userRepo) FindUserByEmailAddress(emailAddress sharedkernelmodel.EmailAddress) (user identitymodel.User, userFound bool, err error) {
	userModels := []pgmodel.UserModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(
			&userModels,
			pgmodel.UserModel{EmailAddress: emailAddress.String()},
		).Error
	}); err != nil {
		return user, userFound, err
	}

	userFound = len(userModels) >= 1
	if !userFound {
		return user, false, nil
	}
	user, err = parseUserModel(userModels[0])
	if err != nil {
		return user, userFound, err
	}
	return user, true, nil
}
