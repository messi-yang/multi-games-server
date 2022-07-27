package gameroomservice

import (
	"errors"
	"math/rand"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/ggol"
	"github.com/google/uuid"
)

var (
	ErrCoordinatesAreInvalid = errors.New("coordinates are not valid")
)

type GameRoomService interface {
	CreateGameRoom(mapSize valueobject.MapSize) *aggregate.GameRoom
	GenerateNextUnitMap(gameId uuid.UUID) error
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

func (gsi *gameRoomServiceImplement) CreateGameRoom(mapSize valueobject.MapSize) *aggregate.GameRoom {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	unitMap := make([][]valueobject.Unit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		unitMap[i] = make([]valueobject.Unit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			unitMap[i][j] = valueobject.NewUnit(rand.Intn(2) == 0, 0)
		}
	}
	gameRoom := aggregate.NewGameRoom()
	gameRoom.UpdateGameMapSize(mapSize)
	gameRoom.UpdateUnitMap(unitMap)
	gsi.gameRoomRepository.Add(gameRoom)

	return &gameRoom
}

func (gsi *gameRoomServiceImplement) GenerateNextUnitMap(gameId uuid.UUID) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	unitMap := gameRoom.GetUnitMap()
	gameOfLiberty, err := ggol.NewGame(&unitMap)
	if err != nil {
		return err
	}
	gameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	nextUnitMap := gameOfLiberty.GenerateNextUnits()
	gsi.gameRoomRepository.UpdateUnitMap(gameId, *nextUnitMap)

	return nil
}

func (gsi *gameRoomServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	areCoordinatesValid := gameRoom.AreCoordinatesValid(coordinates)
	if !areCoordinatesValid {
		return ErrCoordinatesAreInvalid
	}

	for _, coord := range coordinates {
		nextUnit := valueobject.NewUnit(true, 0)
		gsi.gameRoomRepository.UpdateUnit(gameId, coord, nextUnit)
	}

	return nil
}
