package pgrepo

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

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
	userModel := pgmodel.NewUserModel(user)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&userModel).Error
	})
}

func (repo *userRepo) Update(user usermodel.User) error {
	userModel := pgmodel.NewUserModel(user)
	userModel.UpdatedAt = time.Now()
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.UserModel{}).Where(
			"id = ?",
			user.GetId().Uuid(),
		).Select("*").Updates(&userModel).Error
	})
}

func (repo *userRepo) Get(userId globalcommonmodel.UserId) (user usermodel.User, err error) {
	userModel := pgmodel.UserModel{Id: userId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&userModel).Error
	}); err != nil {
		return user, err
	}

	return pgmodel.ParseUserModel(userModel)
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

	user, err := pgmodel.ParseUserModel(userModel)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepo) GetUsersOfIds(userIds []globalcommonmodel.UserId) ([]usermodel.User, error) {
	userModels := []pgmodel.UserModel{}
	userIdDtos := lo.Map(userIds, func(userId globalcommonmodel.UserId, _ int) uuid.UUID {
		return userId.Uuid()
	})

	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(userIdDtos).Find(&userModels).Error
	}); err != nil {
		return nil, err
	}

	users, err := commonutil.MapWithError(userModels, func(_ int, userModel pgmodel.UserModel) (usermodel.User, error) {
		return pgmodel.ParseUserModel(userModel)
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
