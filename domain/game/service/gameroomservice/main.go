package gameroomservice

import (
	"math/rand"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/ggol"
	"github.com/google/uuid"
)

type GameRoomService interface {
	CreateGameRoom(mapSize valueobject.MapSize) *aggregate.GameRoom
	GenerateNextUnitMatrix(gameId uuid.UUID) error
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

	unitMatrix := make([][]valueobject.Unit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		unitMatrix[i] = make([]valueobject.Unit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			unitMatrix[i][j] = valueobject.NewUnit(rand.Intn(2) == 0, 0)
		}
	}
	gameRoom := aggregate.NewGameRoom()
	gameRoom.UpdateGameMapSize(mapSize)
	gameRoom.UpdateUnitMatrix(unitMatrix)
	gsi.gameRoomRepository.Add(gameRoom)

	return &gameRoom
}

func (gsi *gameRoomServiceImplement) GenerateNextUnitMatrix(gameId uuid.UUID) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	unitMatrix := gameRoom.GetUnitMatrix()
	gameOfLiberty, err := ggol.NewGame(&unitMatrix)
	if err != nil {
		return err
	}
	gameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	nextUnitMatrix := gameOfLiberty.GenerateNextUnits()
	gsi.gameRoomRepository.UpdateUnitMatrix(gameId, *nextUnitMatrix)

	return nil
}

func (gsi *gameRoomServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	for _, coord := range coordinates {
		nextUnit := valueobject.NewUnit(true, 0)
		gsi.gameRoomRepository.UpdateUnit(gameId, coord, nextUnit)
	}

	return nil
}
