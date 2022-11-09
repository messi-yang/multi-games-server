package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/infrastructure/memory"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepository    repository.GameRepository
	hardcodedGameIdId valueobject.GameId
}

type gameServiceConfiguration func(service *GameService) error

func NewGameService(cfgs ...gameServiceConfiguration) (*GameService, error) {
	hardcodedGameIdId, _ := uuid.Parse("1a53a474-ebbd-49e4-a2c1-dde5aa5759bc")
	t := &GameService{
		hardcodedGameIdId: valueobject.NewGameId(hardcodedGameIdId),
	}
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

func (gs *GameService) GetAllGameIds() []valueobject.GameId {
	return []valueobject.GameId{gs.hardcodedGameIdId}
}

func (gs *GameService) CreateGame(dimension valueobject.Dimension) (valueobject.GameId, error) {
	unitBlock := make([][]valueobject.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitBlock[i] = make([]valueobject.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitBlock[i][j] = valueobject.NewUnit(false, valueobject.ItemTypeEmpty)
		}
	}
	game := aggregate.NewGame(gs.hardcodedGameIdId, valueobject.NewUnitBlock(unitBlock))
	gs.gameRepository.Add(game)
	return game.GetId(), nil
}

func (service *GameService) GetGame(id valueobject.GameId) (aggregate.Game, error) {
	game, err := service.gameRepository.Get(id)
	if err != nil {
		return aggregate.Game{}, err
	}

	return game, nil
}

func (gs *GameService) AddPlayerToGame(gameId valueobject.GameId, playerId valueobject.PlayerId) (aggregate.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}

	game.AddPlayer(playerId)
	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) RemovePlayerFromGame(gameId valueobject.GameId, playerId valueobject.PlayerId) (aggregate.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}

	game.RemovePlayer(playerId)
	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) AddZoomedAreaToGame(gameId valueobject.GameId, playerId valueobject.PlayerId, area valueobject.Area) (aggregate.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}

	err = game.AddZoomedArea(playerId, area)
	if err != nil {
		return aggregate.Game{}, err
	}

	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) RemoveZoomedAreaFromGame(gameId valueobject.GameId, playerId valueobject.PlayerId) (aggregate.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}

	game.RemoveZoomedArea(playerId)
	gs.gameRepository.Update(gameId, game)

	return game, nil
}

func (gs *GameService) ReviveUnitsInGame(gameId valueobject.GameId, coordinates []valueobject.Coordinate) (aggregate.Game, error) {
	unlocker, err := gs.gameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}
	defer unlocker()

	game, err := gs.gameRepository.Get(gameId)
	if err != nil {
		return aggregate.Game{}, err
	}

	err = game.ReviveUnits(coordinates)
	if err != nil {
		return aggregate.Game{}, err
	}

	gs.gameRepository.Update(gameId, game)

	return game, nil
}
