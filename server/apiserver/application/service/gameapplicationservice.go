package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
)

type GameApplicationService interface {
	GetFirstGameId() (gamemodel.GameId, error)
}

type GameApplicationServe struct {
	GameService gameservice.GameService
}

type gameApplicationServiceConfiguration func(serve *GameApplicationServe) error

func NewGameApplicationService(cfgs ...gameApplicationServiceConfiguration) (*GameApplicationServe, error) {
	serve := &GameApplicationServe{}
	for _, cfg := range cfgs {
		err := cfg(serve)
		if err != nil {
			return nil, err
		}
	}
	return serve, nil
}

func WithGameService() gameApplicationServiceConfiguration {
	gameService, _ := gameservice.NewGameService(gameservice.WithPostgresGameRepository())
	return func(serve *GameApplicationServe) error {
		serve.GameService = gameService
		return nil
	}
}

func (serve *GameApplicationServe) GetFirstGameId() (gamemodel.GameId, error) {
	games, err := serve.GameService.GeAllGames()
	if err != nil {
		return gamemodel.GameId{}, err
	}
	return games[0].GetId(), nil
}
