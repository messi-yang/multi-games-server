package service

import (
	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/repository"
	gameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/infrastructure/gamerepository"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepository repository.GameRepository
}

type gameServiceConfiguration func(service *GameService) error

func NewGameService(cfgs ...gameServiceConfiguration) (*GameService, error) {
	t := &GameService{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithPostgresGameRepository() gameServiceConfiguration {
	return func(service *GameService) error {
		postgresGameRepository, err := gamerepository.NewPostgresGameRepository(gamerepository.WithPostgresClient())
		if err != nil {
			return err
		}
		service.gameRepository = postgresGameRepository
		return nil
	}
}

func (service *GameService) CreateGame(dimension commonValueObject.Dimension) (gameValueObject.GameId, error) {
	unitMatrix := make([][]commonValueObject.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitMatrix[i] = make([]commonValueObject.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitMatrix[i][j] = commonValueObject.NewUnit(false, commonValueObject.ItemTypeEmpty)
		}
	}
	unitBlock := commonValueObject.NewUnitBlock(unitMatrix)

	newGame := aggregate.NewGame(gameValueObject.NewGameId(uuid.New()), unitBlock)
	newGameId, err := service.gameRepository.Add(newGame)
	if err != nil {
		return gameValueObject.GameId{}, err
	}

	return newGameId, nil
}

func (service *GameService) GetGame(gameId gameValueObject.GameId) (aggregate.Game, error) {
	game, err := service.gameRepository.Get(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}

	return game, nil
}

func (service *GameService) GeAllGames() ([]aggregate.Game, error) {
	games, err := service.gameRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return games, nil
}
