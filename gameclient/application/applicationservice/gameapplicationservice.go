package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/google/uuid"
)

type GameApplicationService interface {
	GetFirstGameId() uuid.UUID
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

func (service *gameApplicationServiceImplementation) GetFirstGameId() uuid.UUID {
	games := service.gameDomainService.GetAllGames()
	return games[0].GetId()
}
