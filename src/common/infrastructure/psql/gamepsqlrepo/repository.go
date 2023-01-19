package gamepsqlrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type repo struct {
	gormDb *gorm.DB
}

func New(gormDb *gorm.DB) gamemodel.GameRepo {
	return &repo{gormDb: gormDb}
}

func (m *repo) Get(id gamemodel.GameIdVo) (gamemodel.GameAgg, error) {
	gameModel := GamePsqlModel{Id: id.ToString()}
	result := m.gormDb.First(&gameModel)
	if result.Error != nil {
		return gamemodel.GameAgg{}, result.Error
	}

	return gameModel.ToAggregate(), nil
}

func (m *repo) Update(id gamemodel.GameIdVo, game gamemodel.GameAgg) error {
	return nil
}

func (m *repo) GetAll() ([]gamemodel.GameAgg, error) {
	var gamePsqlModels []GamePsqlModel
	result := m.gormDb.Find(&gamePsqlModels)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := lo.Map(gamePsqlModels, func(model GamePsqlModel, _ int) gamemodel.GameAgg {
		return model.ToAggregate()
	})

	return gameAggregates, nil
}

func (m *repo) Add(game gamemodel.GameAgg) error {
	gameModel := NewGamePsqlModel(game)
	res := m.gormDb.Create(&gameModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (m *repo) ReadLockAccess(gameId gamemodel.GameIdVo) (func(), error) {
	return func() {}, nil
}

func (m *repo) LockAccess(gameId gamemodel.GameIdVo) (func(), error) {
	return func() {}, nil
}
