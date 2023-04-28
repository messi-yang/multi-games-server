package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
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
	db *gorm.DB
}

func NewGamerRepo() (repository gamermodel.Repo, err error) {
	db, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &gamerRepo{db: db}, nil
}

func (repo *gamerRepo) Add(gamer gamermodel.Gamer) error {
	gamerModel := newGamerModel(gamer)
	res := repo.db.Create(&gamerModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *gamerRepo) Get(gamerId commonmodel.GamerId) (gamer gamermodel.Gamer, err error) {
	gamerModel := pgmodel.GamerModel{Id: gamerId.Uuid()}
	result := repo.db.First(&gamerModel)
	if result.Error != nil {
		return gamer, result.Error
	}
	return parseGamerModel(gamerModel), nil
}

func (repo *gamerRepo) GetGamerByUserId(userId sharedkernelmodel.UserId) (gamer gamermodel.Gamer, err error) {
	var gamerModel pgmodel.GamerModel
	result := repo.db.First(&gamerModel, pgmodel.GamerModel{UserId: userId.Uuid()})
	if result.Error != nil {
		return gamer, result.Error
	}
	return parseGamerModel(gamerModel), nil
}

func (repo *gamerRepo) FindGamerByUserId(userId sharedkernelmodel.UserId) (gamer gamermodel.Gamer, gamerFound bool, err error) {
	gamerModels := []pgmodel.GamerModel{}
	result := repo.db.Find(&gamerModels, pgmodel.GamerModel{UserId: userId.Uuid()})
	if result.Error != nil {
		return gamer, gamerFound, result.Error
	}
	gamerFound = result.RowsAffected >= 1
	if !gamerFound {
		return gamer, false, nil
	}
	return parseGamerModel(gamerModels[0]), true, nil
}

func (repo *gamerRepo) GetAll() (gamers []gamermodel.Gamer, err error) {
	var gamerModels []pgmodel.GamerModel
	result := repo.db.Find(&gamerModels)
	if result.Error != nil {
		err = result.Error
		return gamers, err
	}

	gamers = lo.Map(gamerModels, func(gamerModel pgmodel.GamerModel, _ int) gamermodel.Gamer {
		return parseGamerModel(gamerModel)
	})
	return gamers, nil
}
