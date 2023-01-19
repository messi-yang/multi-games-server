package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(mapSize commonmodel.SizeVo) (gamemodel.GameIdVo, error)
}

type gameService struct {
	gameRepo gamemodel.GameRepo
}

func NewGameService(gameRepo gamemodel.GameRepo) GameService {
	return &gameService{gameRepo: gameRepo}
}

func (serve *gameService) CreateGame(mapSize commonmodel.SizeVo) (gamemodel.GameIdVo, error) {
	unitMatrix, _ := tool.RangeMatrix(mapSize.GetWidth(), mapSize.GetHeight(), func(x int, y int) (commonmodel.UnitVo, error) {
		itemId, _ := itemmodel.NewItemIdVo(uuid.Nil.String())
		return commonmodel.NewUnitVo(itemId), nil
	})
	map_ := gamemodel.NewMapVo(unitMatrix)
	newGameId, _ := gamemodel.NewGameIdVo(uuid.New().String())
	newGame := gamemodel.NewGameAgg(newGameId, map_)
	err := serve.gameRepo.Add(newGame)
	if err != nil {
		return gamemodel.GameIdVo{}, err
	}

	return newGameId, nil
}
