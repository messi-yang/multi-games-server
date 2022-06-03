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
	GetGameRoom(gameId uuid.UUID) (*aggregate.GameRoom, error)
	GenerateNextGameUnitMatrix(gameId uuid.UUID) error
	ReviveGameUnit(gameId uuid.UUID, coord valueobject.Coordinate) error
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
	gsi.gameRoomRepository.Create(gameRoom)

	return &gameRoom
}

func (gsi *gameRoomServiceImplement) GetGameRoom(gameId uuid.UUID) (*aggregate.GameRoom, error) {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return nil, err
	}

	return &gameRoom, nil
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
