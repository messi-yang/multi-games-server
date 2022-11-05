package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/google/uuid"
)

type GameApplicationService interface {
	GetFirstSandboxId() (uuid.UUID, error)
}

type gameApplicationServiceImplementation struct {
	sandboxService *sandboxservice.SandboxService
}

type GameApplicationServiceConfiguration struct {
	SandboxService *sandboxservice.SandboxService
}

func NewGameApplicationService(configuration GameApplicationServiceConfiguration) GameApplicationService {
	return &gameApplicationServiceImplementation{
		sandboxService: configuration.SandboxService,
	}
}

func (service *gameApplicationServiceImplementation) GetFirstSandboxId() (uuid.UUID, error) {
	gameId, err := service.sandboxService.GetFirstSandboxId()
	return gameId, err
}
