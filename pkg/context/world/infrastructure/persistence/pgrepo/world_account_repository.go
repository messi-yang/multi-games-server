package pgrepo

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
)

func newWorldAccountModel(user worldaccountmodel.WorldAccount) pgmodel.WorldAccountModel {
	return pgmodel.WorldAccountModel{
		Id:               user.GetId().Uuid(),
		UserId:           user.GetUserId().Uuid(),
		WorldsCount:      user.GetWorldsCount(),
		WorldsCountLimit: user.GetWorldsCountLimit(),
	}
}

func parseWorldAccountModel(worldAccountModel pgmodel.WorldAccountModel) worldaccountmodel.WorldAccount {
	return worldaccountmodel.LoadPlayer(
		worldaccountmodel.NewWorldAccountId(worldAccountModel.Id),
		sharedkernelmodel.NewUserId(worldAccountModel.UserId),
		worldAccountModel.WorldsCount,
		worldAccountModel.WorldsCountLimit,
	)
}

type worldAccountRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewWorldAccountRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository worldaccountmodel.WorldAccountRepo) {
	return &worldAccountRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *worldAccountRepo) Add(worldAccount worldaccountmodel.WorldAccount) error {
	worldAccountModel := newWorldAccountModel(worldAccount)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldAccountModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&worldAccount)
}

func (repo *worldAccountRepo) Update(worldAccount worldaccountmodel.WorldAccount) error {
	worldAccountModel := newWorldAccountModel(worldAccount)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&worldAccountModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&worldAccount)
}

func (repo *worldAccountRepo) Get(worldAccountId worldaccountmodel.WorldAccountId) (worldAccount worldaccountmodel.WorldAccount, err error) {
	worldAccountModel := pgmodel.WorldAccountModel{Id: worldAccountId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldAccountModel).Error
	}); err != nil {
		return worldAccount, err
	}
	return parseWorldAccountModel(worldAccountModel), nil
}

func (repo *worldAccountRepo) GetWorldAccountOfUser(userId sharedkernelmodel.UserId) (worldAccount worldaccountmodel.WorldAccount, err error) {
	var worldAccountModel pgmodel.WorldAccountModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldAccountModel, pgmodel.WorldAccountModel{UserId: userId.Uuid()}).Error
	}); err != nil {
		return worldAccount, err
	}

	return parseWorldAccountModel(worldAccountModel), nil
}

func (repo *worldAccountRepo) GetWorldAccountByUserId(userId sharedkernelmodel.UserId) (*worldaccountmodel.WorldAccount, error) {
	worldAccountModel := pgmodel.WorldAccountModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"user_id = ?",
			userId.Uuid(),
		).First(&worldAccountModel).Error
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return commonutil.ToPointer(parseWorldAccountModel(worldAccountModel)), nil
}

func (repo *worldAccountRepo) GetAll() (worldAccounts []worldaccountmodel.WorldAccount, err error) {
	var worldAccountModels []pgmodel.WorldAccountModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&worldAccountModels).Error
	}); err != nil {
		return worldAccounts, err
	}

	worldAccounts = lo.Map(worldAccountModels, func(worldAccountModel pgmodel.WorldAccountModel, _ int) worldaccountmodel.WorldAccount {
		return parseWorldAccountModel(worldAccountModel)
	})
	return worldAccounts, nil
}
