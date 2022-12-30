package gamepsqlrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"gorm.io/gorm"
)

type repo struct {
	gormDb *gorm.DB
}

func New(gormDb *gorm.DB) gamemodel.GameRepo {
	return &repo{gormDb: gormDb}
}

func (m *repo) Get(id gamemodel.GameId) (gamemodel.Game, error) {
	gameModel := GamePsqlModel{Id: id.ToString()}
	result := m.gormDb.First(&gameModel)
	if result.Error != nil {
		return gamemodel.Game{}, result.Error
	}

	return gameModel.ToAggregate(), nil
}

func (m *repo) Update(id gamemodel.GameId, game gamemodel.Game) error {
	return nil
}

func (m *repo) GetAll() ([]gamemodel.Game, error) {
	var gamePostgresDtos []GamePsqlModel
	result := m.gormDb.Find(&gamePostgresDtos)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := make([]gamemodel.Game, 0)
	for _, gamePostgresDto := range gamePostgresDtos {
		gameAggregates = append(gameAggregates, gamePostgresDto.ToAggregate())
	}

	return gameAggregates, nil
}

func (m *repo) Add(game gamemodel.Game) (gamemodel.GameId, error) {
	gameModel := NewGamePsqlModel(game)
	res := m.gormDb.Create(&gameModel)
	if res.Error != nil {
		return gamemodel.GameId{}, res.Error
	}

	return game.GetId(), nil
}

func (m *repo) ReadLockAccess(gameId gamemodel.GameId) (func(), error) {
	return func() {}, nil
}

func (m *repo) LockAccess(gameId gamemodel.GameId) (func(), error) {
	return func() {}, nil
}
