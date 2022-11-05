package gameservice

import (
	gameModel "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/google/uuid"
)

type GameDomainService interface {
	CreateGame(game sandbox.Sandbox) error

	AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error)
	RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error)

	AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (gameModel.Game, error)
	RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error)

	ReviveUnitsInGame(gameId uuid.UUID, coords []valueobject.Coordinate) (gameModel.Game, error)
}

type gameDomainServiceImplement struct {
	gameRepository gameModel.GameRepository
}

type GameDomainServiceConfiguration struct {
	GameRepository gameModel.GameRepository
}

func NewGameDomainService(coniguration GameDomainServiceConfiguration) GameDomainService {
	return &gameDomainServiceImplement{
		gameRepository: coniguration.GameRepository,
	}
}

func (gsi *gameDomainServiceImplement) CreateGame(sandbox sandbox.Sandbox) error {
	game := gameModel.NewGame(sandbox)
	gsi.gameRepository.Add(game)
	return nil
}

func (gsi *gameDomainServiceImplement) AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gsi.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gsi.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.AddPlayer(playerId)
	gsi.gameRepository.Update(gameId, game)

	return game, nil
}

func (gsi *gameDomainServiceImplement) RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gsi.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gsi.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.RemovePlayer(playerId)
	gsi.gameRepository.Update(gameId, game)

	return game, nil
}

func (gsi *gameDomainServiceImplement) AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (gameModel.Game, error) {
	unlocker, err := gsi.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gsi.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	err = game.AddZoomedArea(playerId, area)
	if err != nil {
		return gameModel.Game{}, err
	}

	gsi.gameRepository.Update(gameId, game)

	return game, nil
}

func (gsi *gameDomainServiceImplement) RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) (gameModel.Game, error) {
	unlocker, err := gsi.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gsi.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	game.RemoveZoomedArea(playerId)
	gsi.gameRepository.Update(gameId, game)

	return game, nil
}

func (gsi *gameDomainServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (gameModel.Game, error) {
	unlocker, err := gsi.gameRepository.LockAccess(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}
	defer unlocker()

	game, err := gsi.gameRepository.Get(gameId)
	if err != nil {
		return gameModel.Game{}, err
	}

	err = game.ReviveUnits(coordinates)
	if err != nil {
		return gameModel.Game{}, err
	}

	gsi.gameRepository.Update(gameId, game)

	return game, nil
}
