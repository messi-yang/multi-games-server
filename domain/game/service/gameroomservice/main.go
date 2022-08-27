package gameroomservice

import (
	"time"

	"github.com/DumDumGeniuss/ggol"
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
	GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, error)
	TickUnitMap(gameId uuid.UUID) (updatedAt time.Time, err error)
	ReviveUnits(gameId uuid.UUID, coords []valueobject.Coordinate) error
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

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return gameRoom, nil
}

func (gsi *gameRoomServiceImplement) GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return valueobject.UnitMap{}, err
	}
	defer rUnlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return valueobject.UnitMap{}, err
	}

	unitMap, err := gameRoom.GetUnitMapByArea(area)
	if err != nil {
		return valueobject.UnitMap{}, err
	}

	return unitMap, nil
}

func (gsi *gameRoomServiceImplement) TickUnitMap(gameId uuid.UUID) (time.Time, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return time.Time{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return time.Time{}, err
	}

	unitMap := gameRoom.GetUnitMap()
	var unitMatrix [][]valueobject.Unit = unitMap.ToUnitMatrix()
	gameOfLiberty, err := ggol.NewGame(&unitMatrix)
	if err != nil {
		return time.Time{}, err
	}
	gameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	nextUnitMatrix := gameOfLiberty.GenerateNextUnits()
	newUnitMap := valueobject.NewUnitMapFromUnitMatrix(*nextUnitMatrix)
	updatedAt, err := gsi.gameRoomRepository.UpdateUnitMap(gameId, newUnitMap)
	if err != nil {
		return time.Time{}, err
	}

	return updatedAt, nil
}

func (gsi *gameRoomServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	resCoords, resUnits, err := gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return err
	}
	gsi.gameRoomRepository.UpdateUnits(gameId, resCoords, resUnits)

	return nil
}
