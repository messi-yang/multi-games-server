package gamedomainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
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
	unitMatrix, _ := tool.RangeMatrix(mapSize.GetWidth(), mapSize.GetHeight(), func(x int, y int) (commonmodel.Unit, error) {
		itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
		return commonmodel.NewUnit(itemId), nil
	})
	unitMap := commonmodel.NewUnitMap(unitMatrix)

	newGameId, _ := gamemodel.NewGameId(uuid.New().String())
	newGame := gamemodel.NewGame(newGameId, unitMap)
	newGameId, err := serve.gameRepo.Add(newGame)
	if err != nil {
		return gamemodel.GameId{}, err
	}

	return newGameId, nil
}
