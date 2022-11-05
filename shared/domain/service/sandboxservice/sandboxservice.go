package sandboxservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox/redis"
	"github.com/google/uuid"
)

type SandboxService struct {
	sandboxRepository sandbox.Repository
}

type SandboxServiceConfiguration func(service *SandboxService) error

func NewSandboxService(cfgs ...SandboxServiceConfiguration) (*SandboxService, error) {
	t := &SandboxService{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithSandboxRedis() SandboxServiceConfiguration {
	sandboxRedis, _ := redis.NewSandboxRedis(redis.WithRedisService())
	return func(service *SandboxService) error {
		service.sandboxRepository = sandboxRedis
		return nil
	}
}

func (service *SandboxService) CreateSandbox(mapSize valueobject.MapSize) (sandbox.Sandbox, error) {
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

func (service *SandboxService) GetSandbox(id uuid.UUID) (sandbox.Sandbox, error) {
	game, err := service.sandboxRepository.Get(id)
	if err != nil {
		return sandbox.Sandbox{}, err
	}

	return game, nil
}

func (service *SandboxService) GetFirstSandboxId() (uuid.UUID, error) {
	gameId, err := service.sandboxRepository.GetFirstGameId()
	return gameId, err
}
