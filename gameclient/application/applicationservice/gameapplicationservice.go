package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
	"github.com/google/uuid"
)

type GameApplicationService interface {
	GetFirstSandboxId() (uuid.UUID, error)
}

type gameApplicationServiceImplementation struct {
	GameService *gameservice.GameService
}

type GameApplicationServiceConfiguration struct {
	GameService *gameservice.GameService
}

func NewGameApplicationService(configuration GameApplicationServiceConfiguration) GameApplicationService {
	return &gameApplicationServiceImplementation{
		GameService: configuration.GameService,
	}
}

func (service *gameApplicationServiceImplementation) GetFirstSandboxId() (uuid.UUID, error) {
	gameId, err := service.GameService.GetFirstSandboxId()
	return gameId, err
}
