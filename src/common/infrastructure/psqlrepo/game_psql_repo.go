package psqlrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type gamePsqlRepo struct {
	gormDb *gorm.DB
}

func NewGamePsqlRepo(gormDb *gorm.DB) gamemodel.GameRepo {
	return &gamePsqlRepo{gormDb: gormDb}
}

func (repo *gamePsqlRepo) Get(id gamemodel.GameIdVo) (gamemodel.GameAgg, error) {
	gameModel := GamePsqlModel{Id: id.ToString()}
	result := repo.gormDb.First(&gameModel)
	if result.Error != nil {
		return gamemodel.GameAgg{}, result.Error
	}

	return gameModel.ToAggregate(), nil
}

func (repo *gamePsqlRepo) Update(id gamemodel.GameIdVo, game gamemodel.GameAgg) error {
	return nil
}

func (repo *gamePsqlRepo) GetAll() ([]gamemodel.GameAgg, error) {
	var gamePsqlModels []GamePsqlModel
	result := repo.gormDb.Find(&gamePsqlModels)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := lo.Map(gamePsqlModels, func(model GamePsqlModel, _ int) gamemodel.GameAgg {
		return model.ToAggregate()
	})

	return gameAggregates, nil
}

func (repo *gamePsqlRepo) Add(game gamemodel.GameAgg) error {
	gameModel := NewGamePsqlModel(game)
	res := repo.gormDb.Create(&gameModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *gamePsqlRepo) ReadLockAccess(gameId gamemodel.GameIdVo) (func(), error) {
	return func() {}, nil
}

func (repo *gamePsqlRepo) LockAccess(gameId gamemodel.GameIdVo) (func(), error) {
	return func() {}, nil
}
