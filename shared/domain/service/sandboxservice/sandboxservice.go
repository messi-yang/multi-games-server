package sandboxservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/sandbox"
	"github.com/google/uuid"
)

type SandboxDomainService interface {
	CreateSandbox(mapSize valueobject.MapSize) (sandbox.Sandbox, error)
	GetSandbox(id uuid.UUID) (sandbox.Sandbox, error)
	GetFirstSandboxId() (uuid.UUID, error)
}

type sandboxDomainServiceImplement struct {
	sandboxRepository sandbox.SandboxRepository
}

type SandboxDomainServiceConfiguration struct {
	SandboxRepository sandbox.SandboxRepository
}

func NewSandboxDomainService(configuration SandboxDomainServiceConfiguration) SandboxDomainService {
	return &sandboxDomainServiceImplement{
		sandboxRepository: configuration.SandboxRepository,
	}
}

func (service *sandboxDomainServiceImplement) CreateSandbox(mapSize valueobject.MapSize) (sandbox.Sandbox, error) {
	unitMap := make([][]valueobject.Unit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		unitMap[i] = make([]valueobject.Unit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			unitMap[i][j] = valueobject.NewUnit(false, uuid.Nil)
		}
	}
	newSandbox := sandbox.NewSandbox(uuid.New(), valueobject.NewUnitMap(unitMap))
	err := service.sandboxRepository.Add(newSandbox)
	if err != nil {
		return sandbox.Sandbox{}, err
	}

	return newSandbox, nil
}

func (service *sandboxDomainServiceImplement) GetSandbox(id uuid.UUID) (sandbox.Sandbox, error) {
	game, err := service.sandboxRepository.Get(id)
	if err != nil {
		return sandbox.Sandbox{}, err
	}

	return game, nil
}

func (service *sandboxDomainServiceImplement) GetFirstSandboxId() (uuid.UUID, error) {
	gameId, err := service.sandboxRepository.GetFirstGameId()
	return gameId, err
}
