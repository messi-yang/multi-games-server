package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
)

type GameApplicationService interface{}

type gameApplicationServe struct {
	gameService gameservice.GameService
}

func NewGameApplicationService(gameService gameservice.GameService) GameApplicationService {
	return &gameApplicationServe{gameService: gameService}
}
