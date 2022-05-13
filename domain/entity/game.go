package entity

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	Id         uuid.UUID
	UnitMatrix [][]valueobject.GameUnit
	MapSize    valueobject.MapSize
}

func NewGame() Game {
	id, _ := uuid.NewUUID()
	return Game{
		Id:         id,
		UnitMatrix: make([][]valueobject.GameUnit, 0),
		MapSize:    valueobject.NewMapSize(0, 0),
	}
}

func (g *Game) GetId() uuid.UUID {
	return g.Id
}

func (g *Game) SetUnitMatrix(unitMatrix [][]valueobject.GameUnit) error {
	g.UnitMatrix = unitMatrix

	return nil
}

func (g *Game) GetUnitMatrix() [][]valueobject.GameUnit {
	return g.UnitMatrix
}

func (g *Game) SetMapSize(mapSize valueobject.MapSize) error {
	g.MapSize = mapSize

	return nil
}

func (g *Game) GetMapSize() valueobject.MapSize {
	return g.MapSize
}
