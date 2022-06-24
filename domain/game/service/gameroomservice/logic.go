package gameroomservice

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/ggol"
)

func gameNextUnitGenerator(
	coord *ggol.Coordinate,
	cell *valueobject.Unit,
	getAdjacentUnit ggol.AdjacentUnitGetter[valueobject.Unit],
) (nextUnit *valueobject.Unit) {
	var aliveAdjacentCellsCount int = 0
	for i := -1; i < 2; i += 1 {
		for j := -1; j < 2; j += 1 {
			if !(i == 0 && j == 0) {
				adjUnit, _ := getAdjacentUnit(coord, &ggol.Coordinate{X: i, Y: j})
				if adjUnit.GetAlive() {
					aliveAdjacentCellsCount += 1
				}
			}
		}
	}
	alive := cell.GetAlive()
	age := cell.GetAge()
	if alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			nextCell := valueobject.NewUnit(false, 0)
			return &nextCell
		} else {
			nextCell := valueobject.NewUnit(alive, age+1)
			return &nextCell
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			nextCell := valueobject.NewUnit(true, 0)
			return &nextCell
		} else {
			return cell
		}
	}
}
