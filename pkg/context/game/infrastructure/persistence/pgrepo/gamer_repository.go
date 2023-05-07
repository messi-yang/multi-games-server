package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
	"github.com/samber/lo"
)

func newGamerModel(user gamermodel.Gamer) pgmodel.GamerModel {
	return pgmodel.GamerModel{
		Id:     user.GetId().Uuid(),
		UserId: user.GetUserId().Uuid(),
	}
}

func parseGamerModel(gamerModel pgmodel.GamerModel) gamermodel.Gamer {
	return gamermodel.NewGamer(
		commonmodel.NewGamerId(gamerModel.Id),
		sharedkernelmodel.NewUserId(gamerModel.UserId),
	)
}

type gamerRepo struct {
	uow pguow.Uow
}

func NewGamerRepo(uow pguow.Uow) (repository gamermodel.Repo) {
	return &gamerRepo{uow: uow}
}

func (repo *gamerRepo) Add(gamer gamermodel.Gamer) error {
	gamerModel := newGamerModel(gamer)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&gamerModel).Error
	})
}

func (repo *gamerRepo) Get(gamerId commonmodel.GamerId) (gamer gamermodel.Gamer, err error) {
	gamerModel := pgmodel.GamerModel{Id: gamerId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&gamerModel).Error
	}); err != nil {
		return gamer, err
	}
	return parseGamerModel(gamerModel), nil
}

func (repo *gamerRepo) GetGamerByUserId(userId sharedkernelmodel.UserId) (gamer gamermodel.Gamer, err error) {
	var gamerModel pgmodel.GamerModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&gamerModel, pgmodel.GamerModel{UserId: userId.Uuid()}).Error
	}); err != nil {
		return gamer, err
	}

	return parseGamerModel(gamerModel), nil
}

func (repo *gamerRepo) FindGamerByUserId(userId sharedkernelmodel.UserId) (gamer gamermodel.Gamer, gamerFound bool, err error) {
	gamerModels := []pgmodel.GamerModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&gamerModels, pgmodel.GamerModel{UserId: userId.Uuid()}).Error
	}); err != nil {
		return gamer, gamerFound, err
	}

	gamerFound = len(gamerModels) >= 1
	if !gamerFound {
		return gamer, false, nil
	}
	return parseGamerModel(gamerModels[0]), true, nil
}

func (repo *gamerRepo) GetAll() (gamers []gamermodel.Gamer, err error) {
	var gamerModels []pgmodel.GamerModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&gamerModels).Error
	}); err != nil {
		return gamers, err
	}

	gamers = lo.Map(gamerModels, func(gamerModel pgmodel.GamerModel, _ int) gamermodel.Gamer {
		return parseGamerModel(gamerModel)
	})
	return gamers, nil
}
