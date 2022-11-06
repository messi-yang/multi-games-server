package gameservice

import (
	gameModel "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/memory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox/redis"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepository    gameModel.Repository
	sandboxRepository sandbox.Repository
}

type gameServiceConfiguration func(service *GameService) error

func NewGameService(cfgs ...gameServiceConfiguration) (*GameService, error) {
	t := &GameService{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithGameMemory() gameServiceConfiguration {
	gameMemory := memory.NewGameMemory()
	return func(service *GameService) error {
		service.gameRepository = gameMemory
		return nil
	}
}

func WithSandboxRedis() gameServiceConfiguration {
	sandboxRedis, _ := redis.NewSandboxRedis(redis.WithRedisService())
	return func(service *GameService) error {
		service.sandboxRepository = sandboxRedis
		return nil
	}
}

func (service *GameService) CreateSandbox(dimension valueobject.Dimension) (sandbox.Sandbox, error) {
	unitBlock := make([][]valueobject.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitBlock[i] = make([]valueobject.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitBlock[i][j] = valueobject.NewUnit(false, uuid.Nil)
		}
	}
	newSandbox := sandbox.NewSandbox(uuid.New(), valueobject.NewUnitBlock(unitBlock))
	err := service.sandboxRepository.Add(newSandbox)
	if err != nil {
		return sandbox.Sandbox{}, err
	}

	return newSandbox, nil
}

func (service *GameService) GetSandbox(id uuid.UUID) (sandbox.Sandbox, error) {
	game, err := service.sandboxRepository.Get(id)
	if err != nil {
		return sandbox.Sandbox{}, err
	}

	return game, nil
}

func (service *GameService) GetFirstSandboxId() (uuid.UUID, error) {
	gameId, err := service.sandboxRepository.GetFirstGameId()
	return gameId, err
}

func (gs *GameService) CreateGame(sandbox sandbox.Sandbox) error {
	game := gameModel.NewGame(sandbox)
	gs.gameRepository.Add(game)
	return nil
}

func (gs *GameService) AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.AddPlayer(playerId)
	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.RemovePlayer(playerId)
	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (gameModel.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	err = game.AddZoomedArea(playerId, area)
	if err != nil {
		return gameModel.Game{}, err
	}

	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.RemoveZoomedArea(playerId)
	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (gameModel.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	err = game.ReviveUnits(coordinates)
	if err != nil {
		return gameModel.Game{}, err
	}

	gs.gameRepository.Update(gameId, game)

	return game, nil
}
