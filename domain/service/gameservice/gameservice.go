package gameservice

import (
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/repository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/ggol"
)

type GameService interface {
	InitializeGame() error
	GenerateNextUnits() error
	ReviveGameUnit(coord *valueobject.Coordinate) error
	GetGameUnitsInArea(area *valueobject.Area) (*[][]*valueobject.GameUnit, error)
	GetGameUnit(coord *valueobject.Coordinate) (*valueobject.GameUnit, error)
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

func (gsi *gameServiceImplement) InjectGameRoomMemoryRepository(gameRoomRepository repository.GameRoomRepository) {
	gsi.gameRoomRepository = gameRoomRepository
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

func (gsi *gameServiceImplement) GenerateNextUnits() error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return err
	}

	gsi.gameOfLiberty.GenerateNextUnits()

	return nil
}

func (gsi *gameServiceImplement) ReviveGameUnit(gameCoordinate *valueobject.Coordinate) error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return err
	}

	coord := convertGameCoordinateToGgolCoordinate(gameCoordinate)

	nextUnit := valueobject.NewGameUnit(true, 0)

	gsi.gameOfLiberty.SetUnit(coord, &nextUnit)

	return nil
}

func (gsi *gameServiceImplement) GetGameUnitsInArea(area *valueobject.Area) (*[][]*valueobject.GameUnit, error) {
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

func (gsi *gameServiceImplement) GetGameUnit(coord *valueobject.Coordinate) (*valueobject.GameUnit, error) {
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
