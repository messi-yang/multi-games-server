package gameservice

import (
	gameModel "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/memory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepository gameModel.Repository
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

func (service *GameService) GetAllGames() []gameModel.Game {
	sandboxId, _ := uuid.Parse("1a53a474-ebbd-49e4-a2c1-dde5aa5759bc")
	sandbox := entity.NewSandbox(sandboxId, valueobject.UnitBlock{})
	game := gameModel.NewGame(sandbox)
	games := make([]gameModel.Game, 0)
	games = append(games, game)
	return games
}

func (gs *GameService) CreateGame(dimension valueobject.Dimension) (uuid.UUID, error) {
	unitBlock := make([][]valueobject.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitBlock[i] = make([]valueobject.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitBlock[i][j] = valueobject.NewUnit(false, uuid.Nil)
		}
	}
	sandboxId, _ := uuid.Parse("1a53a474-ebbd-49e4-a2c1-dde5aa5759bc")
	newSandbox := entity.NewSandbox(sandboxId, valueobject.NewUnitBlock(unitBlock))
	game := gameModel.NewGame(newSandbox)
	gs.gameRepository.Add(game)
	return game.GetId(), nil
}

func (service *GameService) GetGame(id uuid.UUID) (gameModel.Game, error) {
	game, err := service.gameRepository.Get(id)
	if err != nil {
		return gameModel.Game{}, err
	}

	return game, nil
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
