package pgrepo

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newUserModel(user usermodel.User) pgmodel.UserModel {
	return pgmodel.UserModel{
		Id:           user.GetId().Uuid(),
		EmailAddress: user.GetEmailAddress().String(),
		Username:     user.GetUsername().String(),
		CreatedAt:    user.GetCreatedAt(),
		UpdatedAt:    user.GetUpdatedAt(),
	}
}

func parseUserModel(userModel pgmodel.UserModel) (user usermodel.User, err error) {
	emailAddress, err := globalcommonmodel.NewEmailAddress(userModel.EmailAddress)
	if err != nil {
		return user, err
	}
	username, err := globalcommonmodel.NewUsername(userModel.Username)
	if err != nil {
		return user, err
	}
	return usermodel.LoadUser(
		globalcommonmodel.NewUserId(userModel.Id),
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

func NewUserRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository usermodel.UserRepo) {
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

func (repo *userRepo) Update(user usermodel.User) error {
	userModel := newUserModel(user)
	userModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&userModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&user)
}

func (repo *userRepo) Get(userId globalcommonmodel.UserId) (user usermodel.User, err error) {
	userModel := pgmodel.UserModel{Id: userId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&userModel).Error
	}); err != nil {
		return user, err
	}

	return parseUserModel(userModel)
}

func (repo *userRepo) GetUserByEmailAddress(emailAddress globalcommonmodel.EmailAddress) (*usermodel.User, error) {
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

func (repo *userRepo) GetUsersInMap(userIds []globalcommonmodel.UserId) (map[globalcommonmodel.UserId]usermodel.User, error) {
	userModels := []pgmodel.UserModel{}
	userIdDtos := lo.Map(userIds, func(userId globalcommonmodel.UserId, _ int) uuid.UUID {
		return userId.Uuid()
	})
	userMap := make(map[globalcommonmodel.UserId]usermodel.User)

	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(userIdDtos).Find(&userModels).Error
	}); err != nil {
		return nil, err
	}

	for _, userModel := range userModels {
		user, err := parseUserModel(userModel)
		if err != nil {
			return nil, err
		}
		userMap[user.GetId()] = user
	}

	return userMap, nil
}
