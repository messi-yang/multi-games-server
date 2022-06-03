package gameroomservice

import (
	"math/rand"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/repository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/ggol"
	"github.com/google/uuid"
)

type GameRoomService interface {
	CreateGameRoom(mapSize valueobject.MapSize) *aggregate.GameRoom
	GenerateNextGameUnitMatrix(gameId uuid.UUID) error
	ReviveGameUnit(gameId uuid.UUID, coord valueobject.Coordinate) error
}

type gameRoomServiceImplement struct {
	gameRoomRepository repository.GameRoomRepository
	locker             sync.RWMutex
}

var gameRoomService GameRoomService = nil

func NewGameRoomService(gameRoomRepository repository.GameRoomRepository) GameRoomService {
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

	gameUnitMatrix := make([][]valueobject.GameUnit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		gameUnitMatrix[i] = make([]valueobject.GameUnit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			gameUnitMatrix[i][j] = valueobject.NewGameUnit(rand.Intn(2) == 0, 0)
		}
	}
	gameRoom := aggregate.NewGameRoom()
	gameRoom.UpdateGameMapSize(mapSize)
	gameRoom.UpdateGameUnitMatrix(gameUnitMatrix)
	gsi.gameRoomRepository.Add(gameRoom)

	return &gameRoom
}

func (gsi *gameRoomServiceImplement) GenerateNextGameUnitMatrix(gameId uuid.UUID) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameUnitMatrix := gameRoom.GetGameUnitMatrix()
	gameOfLiberty, err := ggol.NewGame(&gameUnitMatrix)
	if err != nil {
		return err
	}
	gameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	nextGameUnitMatrix := gameOfLiberty.GenerateNextUnits()
	gsi.gameRoomRepository.UpdateGameUnitMatrix(gameId, *nextGameUnitMatrix)

	return nil
}

func (gsi *gameRoomServiceImplement) ReviveGameUnit(gameId uuid.UUID, gameCoordinate valueobject.Coordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	nextUnit := valueobject.NewGameUnit(true, 0)
	gsi.gameRoomRepository.UpdateGameUnit(gameId, gameCoordinate, nextUnit)

	return nil
}
