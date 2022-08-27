package gameroomservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameRoomService interface {
	CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error)
	GetAllRooms() []aggregate.GameRoom
	GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error)
	GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, time.Time, error)
	TickUnitMap(gameId uuid.UUID) (time.Time, error)
	ReviveUnits(gameId uuid.UUID, coords []valueobject.Coordinate) (updatedAt time.Time, err error)
}

type gameRoomServiceImplement struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
}

var gameRoomService GameRoomService = nil

func NewGameRoomService(gameRoomRepository gameroomrepository.GameRoomRepository) GameRoomService {
	if gameRoomService == nil {
		gameRoomService = &gameRoomServiceImplement{
			gameRoomRepository: gameRoomRepository,
		}
	}
	return gameRoomService
}

func (gsi *gameRoomServiceImplement) CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error) {
	unitMap := valueobject.NewUnitMap(mapSize)
	game := entity.NewGame(unitMap)
	gameRoom := aggregate.NewGameRoom(game)
	gsi.gameRoomRepository.Add(gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomServiceImplement) GetAllRooms() []aggregate.GameRoom {
	return gsi.gameRoomRepository.GetAll()
}

func (gsi *gameRoomServiceImplement) GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error) {
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

func (gsi *gameRoomServiceImplement) GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, time.Time, error) {
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

func (gsi *gameRoomServiceImplement) TickUnitMap(gameId uuid.UUID) (time.Time, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return time.Time{}, err
	}
	defer unlocker()

	gameRoom, _, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return time.Time{}, err
	}

	newUnitMap, lastTickedAt, err := gameRoom.TickUnitMap()
	if err != nil {
		return time.Time{}, err
	}

	err = gsi.gameRoomRepository.UpdateUnitMap(gameId, newUnitMap)
	if err != nil {
		return time.Time{}, err
	}

	return lastTickedAt, nil
}

func (gsi *gameRoomServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) (time.Time, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return time.Time{}, err
	}
	defer unlocker()

	gameRoom, _, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return time.Time{}, err
	}

	resCoords, resUnits, revivedAt, err := gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return time.Time{}, err
	}
	err = gsi.gameRoomRepository.UpdateUnits(gameId, resCoords, resUnits)
	if err != nil {
		return time.Time{}, err
	}

	return revivedAt, nil
}
