package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	commonpostgresdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/persistence/postgres/dto"
	"gorm.io/gorm"
)

type postgresGameRepository struct {
	postgresClient *gorm.DB
}

func NewPostgresGameRepository(postgresClient *gorm.DB) gamemodel.GameRepository {
	return &postgresGameRepository{postgresClient: postgresClient}
}

func (m *postgresGameRepository) Get(id gamemodel.GameId) (gamemodel.Game, error) {
	gameModel := commonpostgresdto.GamePostgresJsonDto{Id: id.GetId()}
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
	var gamePostgresJsonDtos []commonpostgresdto.GamePostgresJsonDto
	result := m.postgresClient.Find(&gamePostgresJsonDtos)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := make([]gamemodel.Game, 0)
	for _, gamePostgresJsonDto := range gamePostgresJsonDtos {
		gameAggregates = append(gameAggregates, gamePostgresJsonDto.ToAggregate())
	}

	return gameAggregates, nil
}

func (m *postgresGameRepository) Add(game gamemodel.Game) (gamemodel.GameId, error) {
	gameModel := commonpostgresdto.NewGamePostgresJsonDto(game)
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
