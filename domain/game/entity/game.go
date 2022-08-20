package entity

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	id      uuid.UUID
	unitMap [][]valueobject.Unit
	mapSize valueobject.MapSize
}

func NewGame() Game {
	id, _ := uuid.NewUUID()
	return Game{
		id:      id,
		unitMap: make([][]valueobject.Unit, 0),
		mapSize: valueobject.NewMapSize(0, 0),
	}
}

func (g *Game) GetId() uuid.UUID {
	return g.id
}

func (g *Game) GetUnitMap() [][]valueobject.Unit {
	return g.unitMap
}

func (g *Game) SetUnitMap(newUnitMap [][]valueobject.Unit) {
	g.unitMap = newUnitMap
}

func (g *Game) GetMapSize() valueobject.MapSize {
	return g.mapSize
}

func (g *Game) SetMapSize(newMapSize valueobject.MapSize) {
	g.mapSize = newMapSize
}
