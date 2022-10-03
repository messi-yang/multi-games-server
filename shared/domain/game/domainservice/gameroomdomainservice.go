package domainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameRoomDomainService interface {
	CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error)
	GetAllRooms() []aggregate.GameRoom
	GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error)

	AddPlayer(gameId uuid.UUID, player entity.Player) error
	RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error

	AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error
	RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error

	TickUnitMap(gameId uuid.UUID) error
	ReviveUnits(gameId uuid.UUID, coords []valueobject.Coordinate) error
}

type gameRoomDomainServiceImplement struct {
	gameRoomRepository repository.GameRoomRepository
}

func NewGameRoomDomainService(gameRoomRepository repository.GameRoomRepository) GameRoomDomainService {
	return &gameRoomDomainServiceImplement{
		gameRoomRepository: gameRoomRepository,
	}
}

func (gsi *gameRoomDomainServiceImplement) CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error) {
	unitMap := valueobject.NewUnitMap(mapSize)
	game := entity.NewGame(unitMap)
	gameRoom := aggregate.NewGameRoom(game, make(map[uuid.UUID]entity.Player), make(map[uuid.UUID]valueobject.Area))
	gsi.gameRoomRepository.Add(gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) GetAllRooms() []aggregate.GameRoom {
	return gsi.gameRoomRepository.GetAll()
}

func (gsi *gameRoomDomainServiceImplement) GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer rUnlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) AddPlayer(gameId uuid.UUID, player entity.Player) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.AddPlayer(player)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return nil
}

func (gsi *gameRoomDomainServiceImplement) RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.RemovePlayer(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return nil
}

func (gsi *gameRoomDomainServiceImplement) AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.AddZoomedArea(playerId, area)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return nil
}

func (gsi *gameRoomDomainServiceImplement) RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.RemoveZoomedArea(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return nil
}

func (gsi *gameRoomDomainServiceImplement) TickUnitMap(gameId uuid.UUID) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	err = gameRoom.TickUnitMap()
	if err != nil {
		return err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return nil
}

func (gsi *gameRoomDomainServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	err = gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return nil
}
