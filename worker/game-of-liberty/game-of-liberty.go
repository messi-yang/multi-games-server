package game_of_liberty

import (
	"math/rand"
	"time"

	"github.com/DumDumGeniuss/ggol"
)

type CgolCell struct {
	Alive bool `json:"alive"`
}

var game ggol.Game[CgolCell]
var gameUnitsGenerationTicker *time.Ticker

func conwaysGameOfLifeNextUnitGenerator(
	coord *ggol.Coordinate,
	Unit *CgolCell,
	getAdjacentUnit ggol.AdjacentUnitGetter[CgolCell],
) (nextUnit *CgolCell) {
	newUnit := *Unit

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
	if newUnit.Alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			newUnit.Alive = false
			return &newUnit
		} else {
			newUnit.Alive = true
			return &newUnit
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			newUnit.Alive = true
			return &newUnit
		} else {
			newUnit.Alive = false
			return &newUnit
		}
	}
}

func StartGame() {
	if game != nil {
		return
	}
	size := ggol.Size{
		Width:  1000,
		Height: 1000,
	}
	initialCgolCell := CgolCell{
		Alive: false,
	}
	game, _ = ggol.NewGame(&size, &initialCgolCell)
	game.SetNextUnitGenerator(conwaysGameOfLifeNextUnitGenerator)
	game.IterateUnits(func(coord *ggol.Coordinate, _ *CgolCell) {
		game.SetUnit(coord, &CgolCell{Alive: rand.Intn(2) == 0})
	})

	gameUnitsGenerationTicker = time.NewTicker(time.Millisecond * 1000)
	go func() {
		for range gameUnitsGenerationTicker.C {
			game.GenerateNextUnits()
		}
	}()
}

func GetCells() *ggol.Units[CgolCell] {
	if game == nil {
		return nil
	}
	area := ggol.Area{
		From: ggol.Coordinate{X: 0, Y: 0},
		To:   ggol.Coordinate{X: 50, Y: 50},
	}
	units, _ := game.GetUnitsInArea(&area)
	return units
}
