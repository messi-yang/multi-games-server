package gameservice

import (
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/repository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/ggol"
	"github.com/google/uuid"
)

type GameService interface {
	InitializeGame() error
	GenerateNextUnits(gameId uuid.UUID) error
	ReviveGameUnit(gameId uuid.UUID, coord *valueobject.Coordinate) error
}

type gameServiceImplement struct {
	gameOfLiberty      ggol.Game[valueobject.GameUnit]
	gameRoomRepository repository.GameRoomRepository
	locker             sync.RWMutex
}

var gameService GameService = nil

func NewGameService(gameRoomRepository repository.GameRoomRepository) GameService {
	if gameService == nil {
		gameService = &gameServiceImplement{
			gameRoomRepository: gameRoomRepository,
			locker:             sync.RWMutex{},
		}
	}
	return gameService
}

func (gsi *gameServiceImplement) checkIsGameInitialized() error {
	if gsi.gameOfLiberty == nil {
		return &errGameIsNotInitialized{}
	}
	return nil
}

func (gsi *gameServiceImplement) InitializeGame() error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	gameId := config.GetConfig().GetGameId()
	gameRoom, _ := gsi.gameRoomRepository.Get(gameId)

	initialUnit := valueobject.NewGameUnit(false, 0)
	gameSize := gameRoom.GetGameMapSize()
	ggolSize := convertMapSizeToGgolSize(&gameSize)
	newGameOfLiberty, _ := ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	newGameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	gameUnitMatrix := gameRoom.GetGameUnitMatrix()

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &gameUnitMatrix[x][y]
			coord := &ggol.Coordinate{X: x, Y: y}
			newGameOfLiberty.SetUnit(coord, gameFieldUnit)
		}
	}

	gsi.gameOfLiberty = newGameOfLiberty

	return nil
}

func (gsi *gameServiceImplement) GenerateNextUnits(gameId uuid.UUID) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return err
	}

	gsi.gameOfLiberty.GenerateNextUnits()

	gameUnitPointerMatrix := gsi.gameOfLiberty.GetUnits()

	gameUnitMatrix := make([][]valueobject.GameUnit, 0)
	for x := 0; x < len(gameUnitPointerMatrix); x += 1 {
		gameUnitMatrix = append(gameUnitMatrix, make([]valueobject.GameUnit, 0))
		for y := 0; y < len(gameUnitPointerMatrix[x]); y += 1 {
			gameUnitMatrix[x] = append(gameUnitMatrix[x], *gameUnitPointerMatrix[x][y])
		}
	}

	gsi.gameRoomRepository.UpdateGameUnitMatrix(gameId, gameUnitMatrix)

	return nil
}

func (gsi *gameServiceImplement) ReviveGameUnit(gameId uuid.UUID, gameCoordinate *valueobject.Coordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return err
	}

	coord := convertGameCoordinateToGgolCoordinate(gameCoordinate)

	nextUnit := valueobject.NewGameUnit(true, 0)
	gsi.gameOfLiberty.SetUnit(coord, &nextUnit)

	gsi.gameRoomRepository.UpdateGameUnit(gameId, *gameCoordinate, nextUnit)

	return nil
}
