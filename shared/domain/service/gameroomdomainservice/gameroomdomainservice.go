package gameroomdomainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/sandbox"
	"github.com/google/uuid"
)

type GameRoomDomainService interface {
	CreateGameRoom(game sandbox.Sandbox) error

	AddPlayerToGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)
	RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)

	AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (aggregate.GameRoom, error)
	RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)

	ReviveUnitsInGame(gameId uuid.UUID, coords []valueobject.Coordinate) (aggregate.GameRoom, error)
}

type gameRoomDomainServiceImplement struct {
	gameRoomRepository repository.GameRoomRepository
}

type GameRoomDomainServiceConfiguration struct {
	GameRoomRepository repository.GameRoomRepository
}

func NewGameRoomDomainService(coniguration GameRoomDomainServiceConfiguration) GameRoomDomainService {
	return &gameRoomDomainServiceImplement{
		gameRoomRepository: coniguration.GameRoomRepository,
	}
}

func (gsi *gameRoomDomainServiceImplement) CreateGameRoom(game sandbox.Sandbox) error {
	gameRoom := aggregate.NewGameRoom(game)
	gsi.gameRoomRepository.Add(gameRoom)
	return nil
}

func (gsi *gameRoomDomainServiceImplement) AddPlayerToGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.AddPlayer(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.RemovePlayer(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.AddZoomedArea(playerId, area)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.RemoveZoomedArea(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}
