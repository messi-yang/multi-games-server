package gameservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/adapter/persistence"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepository gamemodel.GameRepository
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
		postgresGameRepository, err := persistence.NewPostgresGameRepository(persistence.WithPostgresClient())
		if err != nil {
			return err
		}
		service.gameRepository = postgresGameRepository
		return nil
	}
}

func (service *GameService) CreateGame(dimension gamecommonmodel.Dimension) (gamemodel.GameId, error) {
	unitMatrix := make([][]gamecommonmodel.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitMatrix[i] = make([]gamecommonmodel.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitMatrix[i][j] = gamecommonmodel.NewUnit(false, gamecommonmodel.ItemTypeEmpty)
		}
	}
	unitBlock := gamecommonmodel.NewUnitBlock(unitMatrix)

	newGame := gamemodel.NewGame(gamemodel.NewGameId(uuid.New()), unitBlock)
	newGameId, err := service.gameRepository.Add(newGame)
	if err != nil {
		return gamemodel.GameId{}, err
	}

	return newGameId, nil
}

func (service *GameService) GetGame(gameId gamemodel.GameId) (gamemodel.Game, error) {
	game, err := service.gameRepository.Get(gameId)
	if err != nil {
		return gamemodel.Game{}, err
	}

	return game, nil
}

func (service *GameService) GeAllGames() ([]gamemodel.Game, error) {
	games, err := service.gameRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return games, nil
}
