package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
	"github.com/google/uuid"
)

type GameDomainService interface {
	CreateGame(size commonmodel.Size) (gamemodel.GameId, error)
}

type gameDomainServe struct {
	gameRepo gamemodel.GameRepo
}

func NewGameDomainService(gameRepo gamemodel.GameRepo) GameDomainService {
	return &gameDomainServe{gameRepo: gameRepo}
}

func (serve *gameDomainServe) CreateGame(size commonmodel.Size) (gamemodel.GameId, error) {
	unitMatrix, _ := tool.RangeMatrix(size.GetWidth(), size.GetHeight(), func(x int, y int) (commonmodel.Unit, error) {
		itemId, _ := itemmodel.NewItemId(uuid.Nil.String())
		return commonmodel.NewUnit(itemId), nil
	})
	map_ := commonmodel.NewMap(unitMatrix)

	newGameId, _ := gamemodel.NewGameId(uuid.New().String())
	newGame := gamemodel.NewGame(newGameId, map_)
	newGameId, err := serve.gameRepo.Add(newGame)
	if err != nil {
		return gamemodel.GameId{}, err
	}

	return newGameId, nil
}
