package entity

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	id           uuid.UUID
	unitMap      *valueobject.UnitMap
	lastTickedAt time.Time
}

func NewGame(unitMap *valueobject.UnitMap) Game {
	id, _ := uuid.NewUUID()
	return Game{
		id:           id,
		unitMap:      unitMap,
		lastTickedAt: time.Now(),
	}
}

func NewGameFromExistingEntity(id uuid.UUID, unitMap *valueobject.UnitMap, lastTickedAt time.Time) Game {
	return Game{
		id:           id,
		unitMap:      unitMap,
		lastTickedAt: lastTickedAt,
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

func (g *Game) GetLastTickedAt() time.Time {
	return g.lastTickedAt
}

func (g *Game) SetUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	g.unitMap.SetUnit(coordinate, unit)
}

func (g *Game) SetUnitMap(newUnitMap *valueobject.UnitMap) {
	g.unitMap = newUnitMap
}

func (g *Game) GetUnitMapSize() valueobject.MapSize {
	return g.unitMap.GetMapSize()
}

func (g *Game) SetLastTickedAt(lastTickedAt time.Time) {
	g.lastTickedAt = lastTickedAt
}
