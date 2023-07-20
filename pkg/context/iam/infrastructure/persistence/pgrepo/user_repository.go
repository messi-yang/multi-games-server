package pgrepo

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/identitymodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
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

func (repo *userRepo) Update(user identitymodel.User) error {
	userModel := newUserModel(user)
	userModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&userModel).Error
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

func (repo *userRepo) GetUserByEmailAddress(emailAddress sharedkernelmodel.EmailAddress) (*identitymodel.User, error) {
	userModel := pgmodel.UserModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"email_address = ?",
			emailAddress.String(),
		).First(&userModel).Error
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user, err := parseUserModel(userModel)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
