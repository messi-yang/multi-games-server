package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/google/uuid"
)

type GameApplicationService interface {
	GetFirstGameId() (uuid.UUID, error)
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

func (service *gameApplicationServiceImplementation) GetFirstGameId() (uuid.UUID, error) {
	gameId, err := service.gameDomainService.GetFirstGameId()
	return gameId, err
}
