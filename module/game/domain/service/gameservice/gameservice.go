package gameservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamemodel"
	commonpostgres "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/persistence/postgres"
	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(dimension gamecommonmodel.Dimension) (gamemodel.GameId, error)
	GetGame(gameId gamemodel.GameId) (gamemodel.Game, error)
	GeAllGames() ([]gamemodel.Game, error)
}

type GameServe struct {
	gameRepository gamemodel.GameRepository
}

type gameServiceConfiguration func(serve *GameServe) error

func NewGameService(cfgs ...gameServiceConfiguration) (GameService, error) {
	t := &GameServe{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithPostgresGameRepository() gameServiceConfiguration {
	return func(serve *GameServe) error {
		postgresGameRepository, err := commonpostgres.NewPostgresGameRepository(commonpostgres.WithPostgresClient())
		if err != nil {
			return err
		}
		serve.gameRepository = postgresGameRepository
		return nil
	}
}

func (serve *GameServe) CreateGame(dimension gamecommonmodel.Dimension) (gamemodel.GameId, error) {
	unitMatrix := make([][]gamecommonmodel.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitMatrix[i] = make([]gamecommonmodel.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitMatrix[i][j] = gamecommonmodel.NewUnit(false, gamecommonmodel.ItemTypeEmpty)
		}
	}
	unitBlock := gamecommonmodel.NewUnitBlock(unitMatrix)

	newGame := gamemodel.NewGame(gamemodel.NewGameId(uuid.New()), unitBlock)
	newGameId, err := serve.gameRepository.Add(newGame)
	if err != nil {
		return gamemodel.GameId{}, err
	}

	return newGameId, nil
}

func (serve *GameServe) GetGame(gameId gamemodel.GameId) (gamemodel.Game, error) {
	game, err := serve.gameRepository.Get(gameId)
	if err != nil {
		return gamemodel.Game{}, err
	}

	return game, nil
}

func (serve *GameServe) GeAllGames() ([]gamemodel.Game, error) {
	games, err := serve.gameRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return games, nil
}
