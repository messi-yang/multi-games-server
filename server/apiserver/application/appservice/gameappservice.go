package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/gameservice"
)

type GameAppService interface{}

type gameAppServe struct {
	gameService gameservice.GameService
}

func NewGameAppService(gameService gameservice.GameService) GameAppService {
	return &gameAppServe{gameService: gameService}
}
