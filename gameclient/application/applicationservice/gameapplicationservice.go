package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
)

type GameApplicationService interface {
	GetFirstGame() entity.Game
}

type gameApplicationServiceImplementation struct {
	gameDomainService domainservice.GameDomainService
}

type GameApplicationServiceConfiguration struct {
	GameDomainService domainservice.GameDomainService
}

func NewGameApplicationService(configuration GameApplicationServiceConfiguration) GameApplicationService {
	return &gameApplicationServiceImplementation{
		gameDomainService: configuration.GameDomainService,
	}
}

func (service *gameApplicationServiceImplementation) GetFirstGame() entity.Game {
	games := service.gameDomainService.GetAllGames()
	return games[0]
}
