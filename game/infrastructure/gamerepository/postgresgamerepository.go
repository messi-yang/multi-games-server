package gamerepository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/postgresclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/model"
	"gorm.io/gorm"
)

type postgresGameRepository struct {
	postgresClient *gorm.DB
}

type postgresGameRepositoryConfiguration func(respository *postgresGameRepository) error

func NewPostgresGameRepository(cfgs ...postgresGameRepositoryConfiguration) (repository.GameRepository, error) {
	respository := &postgresGameRepository{}
	for _, cfg := range cfgs {
		err := cfg(respository)
		if err != nil {
			return nil, err
		}
	}
	return respository, nil
}

func WithPostgresClient() postgresGameRepositoryConfiguration {
	return func(repository *postgresGameRepository) error {
		postgresClient, err := postgresclient.NewPostgresClient()
		if err != nil {
			return err
		}
		repository.postgresClient = postgresClient
		return nil
	}
}

func (m *postgresGameRepository) Get(id valueobject.GameId) (aggregate.Game, error) {
	gameModel := model.GameModel{Id: id.GetId()}
	result := m.postgresClient.First(&gameModel)
	if result.Error != nil {
		return aggregate.Game{}, result.Error
	}

	return gameModel.ToAggregate(), nil
}

func (m *postgresGameRepository) Update(id valueobject.GameId, game aggregate.Game) error {
	return nil
}

func (m *postgresGameRepository) GetAll() ([]aggregate.Game, error) {
	var gameModells []model.GameModel
	result := m.postgresClient.Find(&gameModells)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := make([]aggregate.Game, 0)
	for _, gameModell := range gameModells {
		gameAggregates = append(gameAggregates, gameModell.ToAggregate())
	}

	return gameAggregates, nil
}

func (m *postgresGameRepository) Add(game aggregate.Game) (valueobject.GameId, error) {
	gameModel := model.NewGameModel(game)
	res := m.postgresClient.Create(&gameModel)
	if res.Error != nil {
		return valueobject.GameId{}, res.Error
	}

	return game.GetId(), nil
}

func (m *postgresGameRepository) ReadLockAccess(gameId valueobject.GameId) (func(), error) {
	return func() {}, nil
}

func (m *postgresGameRepository) LockAccess(gameId valueobject.GameId) (func(), error) {
	return func() {}, nil
}
