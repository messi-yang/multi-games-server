package entity

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	"github.com/google/uuid"
)

type Sandbox struct {
	id        uuid.UUID
	unitBlock valueobject.UnitBlock
}

func NewSandbox(id uuid.UUID, unitBlock valueobject.UnitBlock) Sandbox {
	return Sandbox{
		id:        id,
		unitBlock: unitBlock,
	}
}

func (g *Sandbox) GetId() uuid.UUID {
	return g.id
}

func (g *Sandbox) GetUnitBlock() valueobject.UnitBlock {
	return g.unitBlock
}

func (g *Sandbox) GetUnit(coordinate valueobject.Coordinate) valueobject.Unit {
	return g.unitBlock.GetUnit(coordinate)
}

func (g *Sandbox) SetUnit(coordinate valueobject.Coordinate, unit valueobject.Unit) {
	g.unitBlock.SetUnit(coordinate, unit)
}

func (g *Sandbox) GetDimension() valueobject.Dimension {
	return g.unitBlock.GetDimension()
}
