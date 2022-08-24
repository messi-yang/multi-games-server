package gameroomservice

import (
	"sync"

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
	GetUnitMapWithArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, error)
	TickUnitMap(gameId uuid.UUID) error
	ReviveUnits(gameId uuid.UUID, coords []valueobject.Coordinate) error
}

type gameRoomServiceImplement struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
	locker             sync.RWMutex
}

var gameRoomService GameRoomService = nil

func NewGameRoomService(gameRoomRepository gameroomrepository.GameRoomRepository) GameRoomService {
	if gameRoomService == nil {
		gameRoomService = &gameRoomServiceImplement{
			gameRoomRepository: gameRoomRepository,
			locker:             sync.RWMutex{},
		}
	}
	return gameRoomService
}

func (gsi *gameRoomServiceImplement) CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error) {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	unitMap := valueobject.NewUnitMap(mapSize)
	game := entity.NewGame(unitMap)
	gameRoom := aggregate.NewGameRoom(game)
	gameRoom.UpdateUnitMap(unitMap)
	gsi.gameRoomRepository.Add(gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomServiceImplement) GetAllRooms() []aggregate.GameRoom {
	return gsi.gameRoomRepository.GetAll()
}

func (gsi *gameRoomServiceImplement) GetUnitMapWithArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, error) {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return valueobject.UnitMap{}, err
	}

	unitMap, err := gameRoom.GetUnitMapWithArea(area)
	if err != nil {
		return valueobject.UnitMap{}, err
	}

	return unitMap, nil
}

func (gsi *gameRoomServiceImplement) TickUnitMap(gameId uuid.UUID) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	unitMap := gameRoom.GetUnitMap()
	var unitMatrix [][]valueobject.Unit = unitMap.ToUnitMatrix()
	gameOfLiberty, err := ggol.NewGame(&unitMatrix)
	if err != nil {
		return err
	}
	gameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	nextUnitMatrix := gameOfLiberty.GenerateNextUnits()
	newUnitMap := valueobject.NewUnitMapFromUnitMatrix(*nextUnitMatrix)
	gsi.gameRoomRepository.UpdateUnitMap(gameId, newUnitMap)

	return nil
}

func (gsi *gameRoomServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

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
