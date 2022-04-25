package gameservice

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/ggol"
)

type GameService interface {
	GenerateNextUnits()
	GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error)
	GetGameSize() *GameSize
}

type gameServiceImplement struct {
	gameOfLiberty ggol.Game[GameUnit]
	gameDAO       gamedao.GameDAO
}

var gameService GameService = nil

func CreateGameService(gameDAO gamedao.GameDAO) (GameService, error) {
	if gameService != nil {
		return nil, &errGameServiceHasBeenCreated{}
	}
	initialUnit := GameUnit{
		Alive: false,
		Age:   0,
	}
	gameSize, _ := gameDAO.GetGameSize()
	ggolSize := convertGameSizeToGgolSize(gameSize)
	newGameOfLiberty, _ := ggol.NewGame(
		ggolSize,
		&initialUnit,
	)

	newGameOfLiberty.SetNextUnitGenerator(gameNextUnitGenerator)

	gameUnitsFromDAO, _ := gameDAO.GetGameUnits()
	gameUnits := convertGameUnitsFromGameDAOToGameUnits(gameUnitsFromDAO)

	for x := 0; x < ggolSize.Width; x += 1 {
		for y := 0; y < ggolSize.Height; y += 1 {
			gameFieldUnit := &(*gameUnits)[x][y]
			coord := &ggol.Coordinate{X: x, Y: y}
			newGameOfLiberty.SetUnit(coord, gameFieldUnit)
		}
	}

	newGameService := &gameServiceImplement{
		gameOfLiberty: newGameOfLiberty,
		gameDAO:       gameDAO,
	}
	gameService = newGameService
	return gameService, nil
}

func (gsi *gameServiceImplement) GenerateNextUnits() {
	gsi.gameOfLiberty.GenerateNextUnits()
}

func (gsi *gameServiceImplement) GetGameUnitsInArea(area *GameArea) (*[][]*GameUnit, error) {
	ggolArea := convertGameAreaToGgolArea(area)
	units, err := gsi.gameOfLiberty.GetUnitsInArea(ggolArea)
	if err != nil {
		return nil, err
	}
	return &units, nil
}

func (gsi *gameServiceImplement) GetGameSize() *GameSize {
	return convertGgolSizeToGameSize(gsi.gameOfLiberty.GetSize())
}
