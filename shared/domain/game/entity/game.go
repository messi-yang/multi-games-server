package entity

import (
	"time"

	"github.com/DumDumGeniuss/ggol"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

func ggolNextUnitGenerator(
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
	if alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			nextCell := valueobject.NewUnit(false)
			return &nextCell
		} else {
			nextCell := valueobject.NewUnit(alive)
			return &nextCell
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			nextCell := valueobject.NewUnit(true)
			return &nextCell
		} else {
			return cell
		}
	}
}

type Game struct {
	id       uuid.UUID
	unitMap  *valueobject.UnitMap
	tickedAt time.Time
	savedAt  time.Time
}

func NewGame(mapSize valueobject.MapSize) Game {
	id, _ := uuid.NewUUID()
	unitMap := valueobject.NewUnitMap(mapSize)
	return Game{
		id:       id,
		unitMap:  unitMap,
		savedAt:  time.Now(),
		tickedAt: time.Now(),
	}
}

func LoadGame(id uuid.UUID, unitMap *valueobject.UnitMap) Game {
	return Game{
		id:       id,
		unitMap:  unitMap,
		savedAt:  time.Now(),
		tickedAt: time.Now(),
	}
}

func (g *Game) GetId() uuid.UUID {
	return g.id
}

func (g *Game) GetUnitMap() *valueobject.UnitMap {
	return g.unitMap
}

func (g *Game) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	return g.unitMap.GetUnit(coordinate)
}

func (g *Game) GetSavedAt() time.Time {
	return g.savedAt
}

func (g *Game) SetSavedAt(savedAt time.Time) {
	g.savedAt = savedAt
}

func (g *Game) GetTickedAt() time.Time {
	return g.tickedAt
}

func (g *Game) SetUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	g.unitMap.SetUnit(coordinate, unit)
}

func (g *Game) GetMapSize() valueobject.MapSize {
	return g.unitMap.GetMapSize()
}

func (g *Game) TickUnitMap() {
	unitMatrix := g.unitMap.ToValueObjectMatrix()
	gameOfLiberty, _ := ggol.NewGame(unitMatrix)
	gameOfLiberty.SetNextUnitGenerator(ggolNextUnitGenerator)
	nextUnitMatrix := gameOfLiberty.GenerateNextUnits()
	g.unitMap = valueobject.NewUnitMapFromUnitMatrix(nextUnitMatrix)
	g.tickedAt = time.Now()
}
