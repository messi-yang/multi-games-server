package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/google/uuid"
)

type GameApplicationService interface {
	GetFirstSandboxId() (uuid.UUID, error)
}

type gameApplicationServiceImplementation struct {
	sandboxDomainService *sandboxservice.SandboxDomainService
}

type GameApplicationServiceConfiguration struct {
	SandboxDomainService *sandboxservice.SandboxDomainService
}

func NewGameApplicationService(configuration GameApplicationServiceConfiguration) GameApplicationService {
	return &gameApplicationServiceImplementation{
		sandboxDomainService: configuration.SandboxDomainService,
	}
}

func (service *gameApplicationServiceImplementation) GetFirstSandboxId() (uuid.UUID, error) {
	gameId, err := service.sandboxDomainService.GetFirstSandboxId()
	return gameId, err
}
