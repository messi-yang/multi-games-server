package game

import (
	"math/rand"

	"github.com/DumDumGeniuss/ggol"
)

type cgolCell struct {
	Alive bool `json:"alive"`
}

func conwaysGameOfLifeNextUnitGenerator(
	coord *ggol.Coordinate,
	cell *cgolCell,
	getAdjacentUnit ggol.AdjacentUnitGetter[cgolCell],
) (nextUnit *cgolCell) {
	nextCell := *cell

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
	if nextCell.Alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			nextCell.Alive = false
			return &nextCell
		} else {
			nextCell.Alive = true
			return &nextCell
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			nextCell.Alive = true
			return &nextCell
		} else {
			nextCell.Alive = false
			return &nextCell
		}
	}
}

func newGame(scale int, blockSize int) ggol.Game[cgolCell] {
	size := ggol.Size{
		Width:  blockSize * scale,
		Height: blockSize * scale,
	}
	initialcgolCell := cgolCell{
		Alive: false,
	}
	game, _ := ggol.NewGame(&size, &initialcgolCell)
	game.SetNextUnitGenerator(conwaysGameOfLifeNextUnitGenerator)
	game.IterateUnits(func(coord *ggol.Coordinate, _ *cgolCell) {
		game.SetUnit(coord, &cgolCell{Alive: rand.Intn(2) == 0})
	})

	return game
}

func getBlockArea(rowIdx int, colIdx int, blockSize int) ggol.Area {
	area := ggol.Area{
		From: ggol.Coordinate{X: blockSize * rowIdx, Y: blockSize * colIdx},
		To:   ggol.Coordinate{X: blockSize*(rowIdx+1) - 1, Y: blockSize*(colIdx+1) - 1},
	}

	return area
}
