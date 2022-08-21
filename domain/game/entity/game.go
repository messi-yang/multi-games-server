package entity

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	id      uuid.UUID
	unitMap [][]valueobject.Unit
}

func NewGame(unitMap [][]valueobject.Unit) Game {
	id, _ := uuid.NewUUID()
	return Game{
		id:      id,
		unitMap: unitMap,
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

func (g *Game) GetUnitMapSize() valueobject.MapSize {
	gameMapSize, _ := valueobject.NewMapSize(len(g.unitMap), len(g.unitMap[0]))
	return gameMapSize
}
