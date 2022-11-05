package sandbox

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type Sandbox struct {
	id      uuid.UUID
	unitMap valueobject.UnitMap
}

func NewSandbox(id uuid.UUID, unitMap valueobject.UnitMap) Sandbox {
	return Sandbox{
		id:      id,
		unitMap: unitMap,
	}
}

func (g *Sandbox) GetId() uuid.UUID {
	return g.id
}

func (g *Sandbox) GetUnitMap() valueobject.UnitMap {
	return g.unitMap
}

func (g *Sandbox) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	return g.unitMap.GetUnit(coordinate)
}

func (g *Sandbox) SetUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	g.unitMap.SetUnit(coordinate, unit)
}

func (g *Sandbox) GetMapSize() valueobject.MapSize {
	return g.unitMap.GetMapSize()
}
