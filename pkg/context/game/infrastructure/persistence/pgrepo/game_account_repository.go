package pgrepo

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gameaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
)

type gameAccountRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewGameAccountRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository gameaccountmodel.GameAccountRepo) {
	return &gameAccountRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *gameAccountRepo) Add(gameAccount gameaccountmodel.GameAccount) error {
	gameaccountmodel := pgmodel.NewGameAccountModel(gameAccount)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&gameaccountmodel).Error
	})
}

func (repo *gameAccountRepo) Update(gameAccount gameaccountmodel.GameAccount) error {
	gameaccountmodel := pgmodel.NewGameAccountModel(gameAccount)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.GameAccountModel{}).Where(
			"id = ?",
			gameAccount.GetId().Uuid(),
		).Select("*").Updates(gameaccountmodel).Error
	})
}

func (repo *gameAccountRepo) Get(gameAccountId gameaccountmodel.GameAccountId) (gameAccount gameaccountmodel.GameAccount, err error) {
	gameaccountmodel := pgmodel.GameAccountModel{Id: gameAccountId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&gameaccountmodel).Error
	}); err != nil {
		return gameAccount, err
	}
	return pgmodel.ParseGameAccountModel(gameaccountmodel), nil
}

func (repo *gameAccountRepo) GetGameAccountOfUser(userId globalcommonmodel.UserId) (gameAccount gameaccountmodel.GameAccount, err error) {
	var gameaccountmodel pgmodel.GameAccountModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&gameaccountmodel, pgmodel.GameAccountModel{UserId: userId.Uuid()}).Error
	}); err != nil {
		return gameAccount, err
	}

	return pgmodel.ParseGameAccountModel(gameaccountmodel), nil
}

func (repo *gameAccountRepo) GetGameAccountByUserId(userId globalcommonmodel.UserId) (*gameaccountmodel.GameAccount, error) {
	gameaccountmodel := pgmodel.GameAccountModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"user_id = ?",
			userId.Uuid(),
		).First(&gameaccountmodel).Error
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return commonutil.ToPointer(pgmodel.ParseGameAccountModel(gameaccountmodel)), nil
}

func (repo *gameAccountRepo) GetAll() (gameAccounts []gameaccountmodel.GameAccount, err error) {
	var gameaccountmodels []pgmodel.GameAccountModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&gameaccountmodels).Error
	}); err != nil {
		return gameAccounts, err
	}

	gameAccounts = lo.Map(gameaccountmodels, func(gameaccountmodel pgmodel.GameAccountModel, _ int) gameaccountmodel.GameAccount {
		return pgmodel.ParseGameAccountModel(gameaccountmodel)
	})
	return gameAccounts, nil
}
