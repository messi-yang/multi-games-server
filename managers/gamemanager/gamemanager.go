package gamemanager

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/ggol"
)

type GameUnit struct {
	Alive bool
}

func gameOfLifeNextUnitGenerator(
	coord *ggol.Coordinate,
	cell *GameUnit,
	getAdjacentUnit ggol.AdjacentUnitGetter[GameUnit],
) (nextUnit *GameUnit) {
	var aliveAdjacentCellsCount int = 0
	for i := -1; i < 2; i += 1 {
		for j := -1; j < 2; j += 1 {
			if !(i == 0 && j == 0) {
				adjUnit, _ := getAdjacentUnit(coord, &ggol.Coordinate{X: i, Y: j})
				if adjUnit.Alive {
					aliveAdjacentCellsCount += 1
				}
			}
		}
	}
	if cell.Alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			nextCell := *cell
			nextCell.Alive = false
			return &nextCell
		} else {
			return cell
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			nextCell := *cell
			nextCell.Alive = true
			return &nextCell
		} else {
			return cell
		}
	}
}

type GameManager struct {
	Game ggol.Game[GameUnit]
}

var gameManager GameManager

func GetUnits(fromX int, fromY int, toX int, toY int) ([][]*GameUnit, error) {
	units, err := gameManager.Game.GetUnitsInArea(&ggol.Area{
		From: ggol.Coordinate{X: fromX, Y: fromY},
		To:   ggol.Coordinate{X: toX, Y: toY},
	})
	if err != nil {
		return nil, err
	}
	return units, nil
}

func TickGame() {
	gameManager.Game.GenerateNextUnits()
}

func InitializeGameManager() {
	size := ggol.Size{
		Width:  config.Config.GAME_SIZE,
		Height: config.Config.GAME_SIZE,
	}
	initialGameUnit := GameUnit{
		Alive: false,
	}
	newGame, err := ggol.NewGame(&size, &initialGameUnit)
	if err != nil {
		panic(err)
	}
	newGame.SetNextUnitGenerator(gameOfLifeNextUnitGenerator)
	newGame.IterateUnits(func(coord *ggol.Coordinate, _ *GameUnit) {
		newGame.SetUnit(coord, &GameUnit{Alive: rand.Intn(2) == 0})
	})
	gameManager.Game = newGame
}
