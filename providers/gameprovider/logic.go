package gameprovider

import "github.com/DumDumGeniuss/ggol"

func gameNextUnitGenerator(
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
			nextCell.Age = 0
			return &nextCell
		} else {
			nextCell := *cell
			nextCell.Age += 1
			return &nextCell
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			nextCell := *cell
			nextCell.Alive = true
			nextCell.Age = 0
			return &nextCell
		} else {
			return cell
		}
	}
}
