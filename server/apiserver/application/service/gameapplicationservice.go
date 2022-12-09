package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
)

type GameApplicationService interface {
	GetFirstGameId() (gamemodel.GameId, error)
}

type gameApplicationServe struct {
	gameService gameservice.GameService
}

func NewGameApplicationService(gameService gameservice.GameService) *gameApplicationServe {
	return &gameApplicationServe{gameService: gameService}
}

func (serve *gameApplicationServe) GetFirstGameId() (gamemodel.GameId, error) {
	games, err := serve.gameService.GeAllGames()
	if err != nil {
		return gamemodel.GameId{}, err
	}
	return games[0].GetId(), nil
}
