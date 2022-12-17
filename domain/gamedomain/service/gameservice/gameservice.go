package gameservice

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(dimension gamecommonmodel.Dimension) (gamemodel.GameId, error)
}

type GameServe struct {
	gameRepository gamemodel.GameRepository
}

func NewGameService(gameRepository gamemodel.GameRepository) GameService {
	return &GameServe{gameRepository: gameRepository}
}

func (serve *GameServe) CreateGame(dimension gamecommonmodel.Dimension) (gamemodel.GameId, error) {
	unitMatrix := make([][]gamecommonmodel.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitMatrix[i] = make([]gamecommonmodel.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitMatrix[i][j] = gamecommonmodel.NewUnit(false, itemmodel.NewItemId(uuid.Nil))
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
