package gameroomservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type Service interface {
	CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error)
	GetAllRooms() []aggregate.GameRoom
	GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error)
	GetUnitMap(gameId uuid.UUID) (valueobject.UnitMap, time.Time, error)
	GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, time.Time, error)
	TickUnitMap(gameId uuid.UUID) (time.Time, error)
	ReviveUnits(gameId uuid.UUID, coords []valueobject.Coordinate) (time.Time, error)
}

type serviceImplement struct {
	gameRoomRepository gameroomrepository.Repository
}

var serviceInstance Service = nil

func NewService(gameRoomRepository gameroomrepository.Repository) Service {
	if serviceInstance == nil {
		serviceInstance = &serviceImplement{
			gameRoomRepository: gameRoomRepository,
		}
	}
	return serviceInstance
}

func (gsi *serviceImplement) CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error) {
	unitMap := valueobject.NewUnitMap(mapSize)
	game := entity.NewGame(unitMap)
	gameRoom := aggregate.NewGameRoom(game)
	gsi.gameRoomRepository.Add(gameRoom)

	return gameRoom, nil
}

func (gsi *serviceImplement) GetAllRooms() []aggregate.GameRoom {
	return gsi.gameRoomRepository.GetAll()
}

func (gsi *serviceImplement) GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer rUnlocker()

	gameRoom, _, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return gameRoom, nil
}

func (gsi *serviceImplement) GetUnitMap(gameId uuid.UUID) (valueobject.UnitMap, time.Time, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return valueobject.UnitMap{}, time.Time{}, err
	}
	defer rUnlocker()

	gameRoom, receivedAt, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return valueobject.UnitMap{}, time.Time{}, err
	}

	unitMap := gameRoom.GetUnitMap()
	if err != nil {
		return valueobject.UnitMap{}, time.Time{}, err
	}

	return unitMap, receivedAt, nil
}

func (gsi *serviceImplement) GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, time.Time, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return valueobject.UnitMap{}, time.Time{}, err
	}
	defer rUnlocker()

	gameRoom, receivedAt, err := gsi.gameRoomRepository.Get(gameId)

	if err != nil {
		return valueobject.UnitMap{}, time.Time{}, err
	}

	unitMap, err := gameRoom.GetUnitMapByArea(area)
	if err != nil {
		return valueobject.UnitMap{}, time.Time{}, err
	}

	return unitMap, receivedAt, nil
}

func (gsi *serviceImplement) TickUnitMap(gameId uuid.UUID) (time.Time, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return time.Time{}, err
	}
	defer unlocker()

	gameRoom, _, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return time.Time{}, err
	}

	tickedAt, err := gameRoom.TickUnitMap()
	if err != nil {
		return time.Time{}, err
	}

	updatedUnitMap := gameRoom.GetUnitMap()

	err = gsi.gameRoomRepository.UpdateUnitMap(gameId, updatedUnitMap)
	if err != nil {
		return time.Time{}, err
	}

	err = gsi.gameRoomRepository.UpdateLastTickedAt(gameId, tickedAt)
	if err != nil {
		return time.Time{}, err
	}

	return tickedAt, nil
}

func (gsi *serviceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) (time.Time, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return time.Time{}, err
	}
	defer unlocker()

	gameRoom, _, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return time.Time{}, err
	}

	revivedAt, err := gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return time.Time{}, err
	}

	updatedUnits, err := gameRoom.GetUnitsWithCoordinates(coordinates)
	if err != nil {
		return time.Time{}, err
	}

	err = gsi.gameRoomRepository.UpdateUnits(gameId, coordinates, updatedUnits)
	if err != nil {
		return time.Time{}, err
	}

	return revivedAt, nil
}
