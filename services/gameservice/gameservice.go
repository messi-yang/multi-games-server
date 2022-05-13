package gameservice

import (
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory"
	"github.com/DumDumGeniuss/ggol"
)

type GameService interface {
	InjectGameRoomMemoryRepository(memory.GameRoomMemoryRepository)
	InitializeGame() error
	GenerateNextUnits() error
	ReviveGameUnit(coord *GameCoordinate) error
	GetGameUnitsInArea(area *GameArea) (*[][]*valueobject.GameUnit, error)
	GetMapSize() (*valueobject.MapSize, error)
	GetGameUnit(coord *GameCoordinate) (*valueobject.GameUnit, error)
}

type gameServiceImplement struct {
	gameOfLiberty            ggol.Game[valueobject.GameUnit]
	gameRoomMemoryRepository memory.GameRoomMemoryRepository
	locker                   sync.RWMutex
}

var gameService GameService = nil

func GetGameService() GameService {
	if gameService == nil {
		gameService = &gameServiceImplement{
			locker: sync.RWMutex{},
		}
	}
	return gameService
}

func (gsi *gameServiceImplement) checkGameRoomMemoryRepositoryDependency() error {
	if gsi.gameRoomMemoryRepository == nil {
		return &errMissingGameRoomMemoryRepositoryDependency{}
	}
	return nil
}

func (gsi *gameServiceImplement) checkIsGameInitialized() error {
	if gsi.gameOfLiberty == nil {
		return &errGameIsNotInitialized{}
	}
	return nil
}

func (gsi *gameServiceImplement) InjectGameRoomMemoryRepository(gameRoomMemoryRepository memory.GameRoomMemoryRepository) {
	gsi.gameRoomMemoryRepository = gameRoomMemoryRepository
}

func (gsi *gameServiceImplement) InitializeGame() error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkGameRoomMemoryRepositoryDependency(); err != nil {
		return err
	}

	initialUnit := valueobject.GameUnit{
		Alive: false,
		Age:   0,
	}
	gameSize := gsi.gameRoomMemoryRepository.GetMapSize()
	ggolSize := convertMapSizeToGgolSize(gameSize)
	newGameOfLiberty, _ := ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	newGameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	gameUnitMatrix := gsi.gameRoomMemoryRepository.GetGameUnitMatrix()

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &(*gameUnitMatrix)[x][y]
			coord := &ggol.Coordinate{X: x, Y: y}
			newGameOfLiberty.SetUnit(coord, gameFieldUnit)
		}
	}

	gsi.gameOfLiberty = newGameOfLiberty

	return nil
}

func (gsi *gameServiceImplement) GenerateNextUnits() error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return err
	}

	gsi.gameOfLiberty.GenerateNextUnits()

	return nil
}

func (gsi *gameServiceImplement) ReviveGameUnit(gameCoordinate *GameCoordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return err
	}

	coord := convertGameCoordinateToGgolCoordinate(gameCoordinate)
	unit, err := gsi.gameOfLiberty.GetUnit(coord)
	if err != nil {
		return err
	}

	nextUnit := *unit
	nextUnit.Alive = true

	gsi.gameOfLiberty.SetUnit(coord, &nextUnit)

	return nil
}

func (gsi *gameServiceImplement) GetGameUnitsInArea(area *GameArea) (*[][]*valueobject.GameUnit, error) {
	gsi.locker.RLock()
	defer gsi.locker.RUnlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return nil, err
	}

	ggolArea := convertGameAreaToGgolArea(area)
	units, err := gsi.gameOfLiberty.GetUnitsInArea(ggolArea)
	if err != nil {
		return nil, err
	}
	return &units, nil
}

func (gsi *gameServiceImplement) GetMapSize() (*valueobject.MapSize, error) {
	gsi.locker.RLock()
	defer gsi.locker.RUnlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return nil, err
	}

	return convertGgolSizeToMapSize(gsi.gameOfLiberty.GetSize()), nil
}

func (gsi *gameServiceImplement) GetGameUnit(coord *GameCoordinate) (*valueobject.GameUnit, error) {
	gsi.locker.RLock()
	defer gsi.locker.RUnlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return nil, err
	}

	ggolCoord := convertGameCoordinateToGgolCoordinate(coord)
	gameUnit, err := gsi.gameOfLiberty.GetUnit(ggolCoord)
	if err != nil {
		return nil, err
	}

	return gameUnit, nil
}
