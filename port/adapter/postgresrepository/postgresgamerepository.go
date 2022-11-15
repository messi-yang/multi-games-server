package postgresrepository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/postgresclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/postgresrepository/postgresdto"
	"gorm.io/gorm"
)

type postgresGameRepository struct {
	postgresClient *gorm.DB
}

type postgresGameRepositoryConfiguration func(respository *postgresGameRepository) error

func NewPostgresGameRepository(cfgs ...postgresGameRepositoryConfiguration) (gamemodel.GameRepository, error) {
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

func (m *postgresGameRepository) Get(id gamemodel.GameId) (gamemodel.Game, error) {
	gameModel := postgresdto.GamePostgresUiDto{Id: id.GetId()}
	result := m.postgresClient.First(&gameModel)
	if result.Error != nil {
		return gamemodel.Game{}, result.Error
	}

	return gameModel.ToAggregate(), nil
}

func (m *postgresGameRepository) Update(id gamemodel.GameId, game gamemodel.Game) error {
	return nil
}

func (m *postgresGameRepository) GetAll() ([]gamemodel.Game, error) {
	var gamePostgresUiDtos []postgresdto.GamePostgresUiDto
	result := m.postgresClient.Find(&gamePostgresUiDtos)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := make([]gamemodel.Game, 0)
	for _, gamePostgresUiDto := range gamePostgresUiDtos {
		gameAggregates = append(gameAggregates, gamePostgresUiDto.ToAggregate())
	}

	return gameAggregates, nil
}

func (m *postgresGameRepository) Add(game gamemodel.Game) (gamemodel.GameId, error) {
	gameModel := postgresdto.NewGamePostgresUiDto(game)
	res := m.postgresClient.Create(&gameModel)
	if res.Error != nil {
		return gamemodel.GameId{}, res.Error
	}

	return game.GetId(), nil
}

func (m *postgresGameRepository) ReadLockAccess(gameId gamemodel.GameId) (func(), error) {
	return func() {}, nil
}

func (m *postgresGameRepository) LockAccess(gameId gamemodel.GameId) (func(), error) {
	return func() {}, nil
}
