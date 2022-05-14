package entity

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	id         uuid.UUID
	unitMatrix [][]valueobject.GameUnit
	mapSize    valueobject.MapSize
}

func NewGame() Game {
	id, _ := uuid.NewUUID()
	return Game{
		id:         id,
		unitMatrix: make([][]valueobject.GameUnit, 0),
		mapSize:    valueobject.NewMapSize(0, 0),
	}
}

func (g *Game) GetId() uuid.UUID {
	return g.id
}

func (g *Game) SetUnitMatrix(unitMatrix [][]valueobject.GameUnit) error {
	g.unitMatrix = unitMatrix

	return nil
}

func (g *Game) GetUnitMatrix() [][]valueobject.GameUnit {
	return g.unitMatrix
}

func (g *Game) SetMapSize(mapSize valueobject.MapSize) error {
	g.mapSize = mapSize

	return nil
}

func (g *Game) GetMapSize() valueobject.MapSize {
	return g.mapSize
}
