package gameservice

import (
	gameModel "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/memory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepository gameModel.Repository
}

type GameServiceConfiguration func(os *GameService) error

func NewGameService(cfgs ...GameServiceConfiguration) (*GameService, error) {
	t := &GameService{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithGameMemory() GameServiceConfiguration {
	gameMemory := memory.NewGameMemory()
	return func(service *GameService) error {
		service.gameRepository = gameMemory
		return nil
	}
}

func (gds GameService) CreateGame(sandbox sandbox.Sandbox) error {
	game := gameModel.NewGame(sandbox)
	gds.gameRepository.Add(game)
	return nil
}

func (gds GameService) AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gds.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gds.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.AddPlayer(playerId)
	gds.gameRepository.Update(gameId, game)

	return game, nil
}

func (gds GameService) RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gds.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gds.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.RemovePlayer(playerId)
	gds.gameRepository.Update(gameId, game)

	return game, nil
}

func (gds GameService) AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (gameModel.Game, error) {
	unlocker, err := gds.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gds.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	err = game.AddZoomedArea(playerId, area)
	if err != nil {
		return gameModel.Game{}, err
	}

	gds.gameRepository.Update(gameId, game)

	return game, nil
}

func (gds GameService) RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gds.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gds.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.RemoveZoomedArea(playerId)
	gds.gameRepository.Update(gameId, game)

	return game, nil
}

func (gds GameService) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (gameModel.Game, error) {
	unlocker, err := gds.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gds.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	err = game.ReviveUnits(coordinates)
	if err != nil {
		return gameModel.Game{}, err
	}

	gds.gameRepository.Update(gameId, game)

	return game, nil
}
