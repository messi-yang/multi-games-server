package entity

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	id         uuid.UUID
	unitMap    *valueobject.UnitMap
	tickedAt   time.Time
	savedAt    time.Time
	tickPeriod int64
}

func NewGame(mapSize valueobject.MapSize, tickPeriod int64) Game {
	id, _ := uuid.NewUUID()
	unitMap := valueobject.NewUnitMap(mapSize)
	return Game{
		id:         id,
		unitMap:    unitMap,
		savedAt:    time.Now(),
		tickedAt:   time.Now(),
		tickPeriod: tickPeriod,
	}
}

func LoadGame(id uuid.UUID, unitMap *valueobject.UnitMap, tickPeriod int64) Game {
	return Game{
		id:         id,
		unitMap:    unitMap,
		savedAt:    time.Now(),
		tickedAt:   time.Now(),
		tickPeriod: tickPeriod,
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

func (g *Game) SetUnitMap(newUnitMap *valueobject.UnitMap) {
	g.unitMap = newUnitMap
}

func (g *Game) GetMapSize() valueobject.MapSize {
	return g.unitMap.GetMapSize()
}

func (g *Game) SetTickedAt(tickedAt time.Time) {
	g.tickedAt = tickedAt
}
