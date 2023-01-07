package gamedomainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
)

type Service interface {
	CreateGame(mapSize commonmodel.MapSize) (gamemodel.GameId, error)
}

type servce struct {
	gameRepo gamemodel.GameRepo
}

func New(gameRepo gamemodel.GameRepo) Service {
	return &servce{gameRepo: gameRepo}
}

func (serve *servce) CreateGame(mapSize commonmodel.MapSize) (gamemodel.GameId, error) {
	mapUnitMatrix := make([][]commonmodel.MapUnit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		mapUnitMatrix[i] = make([]commonmodel.MapUnit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
			mapUnitMatrix[i][j] = commonmodel.NewMapUnit(itemId)
		}
	}
	unitMap := commonmodel.NewUnitMap(mapUnitMatrix)

	newGameId, _ := gamemodel.NewGameId(uuid.New().String())
	newGame := gamemodel.NewGame(newGameId, unitMap)
	newGameId, err := serve.gameRepo.Add(newGame)
	if err != nil {
		return gamemodel.GameId{}, err
	}

	return newGameId, nil
}
