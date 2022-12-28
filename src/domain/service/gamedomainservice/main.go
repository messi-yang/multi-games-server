package gamedomainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
)

type Service interface {
	CreateGame(dimension commonmodel.Dimension) (gamemodel.GameId, error)
}

type servce struct {
	gameRepository gamemodel.GameRepository
}

func New(gameRepository gamemodel.GameRepository) Service {
	return &servce{gameRepository: gameRepository}
}

func (serve *servce) CreateGame(dimension commonmodel.Dimension) (gamemodel.GameId, error) {
	unitMatrix := make([][]commonmodel.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitMatrix[i] = make([]commonmodel.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
			unitMatrix[i][j] = commonmodel.NewUnit(itemId)
		}
	}
	unitBlock := commonmodel.NewUnitBlock(unitMatrix)

	newGame := gamemodel.NewGame(gamemodel.NewGameId(uuid.New()), unitBlock)
	newGameId, err := serve.gameRepository.Add(newGame)
	if err != nil {
		return gamemodel.GameId{}, err
	}

	return newGameId, nil
}
