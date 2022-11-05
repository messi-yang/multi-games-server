package gameroomdomainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/google/uuid"
)

type GameRoomDomainService interface {
	CreateGameRoom(game sandbox.Sandbox) error

	AddPlayerToGameRoom(gameId uuid.UUID, playerId uuid.UUID) (game.GameRoom, error)
	RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (game.GameRoom, error)

	AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (game.GameRoom, error)
	RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (game.GameRoom, error)

	ReviveUnitsInGame(gameId uuid.UUID, coords []valueobject.Coordinate) (game.GameRoom, error)
}

type gameRoomDomainServiceImplement struct {
	gameRoomRepository game.GameRoomRepository
}

type GameRoomDomainServiceConfiguration struct {
	GameRoomRepository game.GameRoomRepository
}

func NewGameRoomDomainService(coniguration GameRoomDomainServiceConfiguration) GameRoomDomainService {
	return &gameRoomDomainServiceImplement{
		gameRoomRepository: coniguration.GameRoomRepository,
	}
}

func (gsi *gameRoomDomainServiceImplement) CreateGameRoom(sandbox sandbox.Sandbox) error {
	gameRoom := game.NewGameRoom(sandbox)
	gsi.gameRoomRepository.Add(gameRoom)
	return nil
}

func (gsi *gameRoomDomainServiceImplement) AddPlayerToGameRoom(gameId uuid.UUID, playerId uuid.UUID) (game.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}

	gameRoom.AddPlayer(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (game.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}

	gameRoom.RemovePlayer(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (game.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}

	err = gameRoom.AddZoomedArea(playerId, area)
	if err != nil {
		return game.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (game.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}

	gameRoom.RemoveZoomedArea(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (game.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return game.GameRoom{}, err
	}

	err = gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return game.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}
