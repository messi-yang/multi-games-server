package gameservice

import (
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameService interface {
	InjectGameDAO(gameDAO gamedao.GameDAO)
	InitializeGame() error
	GenerateNextUnits() error
	ReviveGameUnit(coord *GameCoordinate) error
	GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error)
	GetGameSize() (*GameSize, error)
}

type gameServiceImplement struct {
	gameOfLiberty ggol.Game[GameUnit]
	gameDAO       gamedao.GameDAO
	locker        sync.RWMutex
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

func (gsi *gameServiceImplement) checkGameDAODependency() error {
	if gsi.gameDAO == nil {
		return &errMissingGameDAODependency{}
	}
	return nil
}

func (gsi *gameServiceImplement) checkIsGameInitialized() error {
	if gsi.gameOfLiberty == nil {
		return &errGameIsNotInitialized{}
	}
	return nil
}

func (gsi *gameServiceImplement) InjectGameDAO(gameDAO gamedao.GameDAO) {
	gsi.gameDAO = gameDAO
}

func (gsi *gameServiceImplement) InitializeGame() error {
	gsi.locker.Lock()
	defer gsi.locker.Unlock()

	if err := gsi.checkGameDAODependency(); err != nil {
		return err
	}

	initialUnit := GameUnit{
		Alive: false,
		Age:   0,
	}
	gameSize, _ := gsi.gameDAO.GetGameSize()
	ggolSize := convertGameSizeToGgolSize(gameSize)
	newGameOfLiberty, _ := ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	newGameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	gameUnitsFromDAO, _ := gsi.gameDAO.GetGameUnits()
	gameUnits := convertGameUnitsFromGameDAOToGameUnits(gameUnitsFromDAO)

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &(*gameUnits)[x][y]
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

func (gsi *gameServiceImplement) GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error) {
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

func (gsi *gameServiceImplement) GetGameSize() (*GameSize, error) {
	gsi.locker.RLock()
	defer gsi.locker.RUnlock()

	if err := gsi.checkIsGameInitialized(); err != nil {
		return nil, err
	}

	return convertGgolSizeToGameSize(gsi.gameOfLiberty.GetSize()), nil
}
