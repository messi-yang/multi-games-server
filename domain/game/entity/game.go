package entity

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

func getUnitMapSize(unitMap valueobject.UnitMap) valueobject.MapSize {
	gameMapSize, _ := valueobject.NewMapSize(len(unitMap), len(unitMap[0]))
	return gameMapSize
}

type Game struct {
	id          uuid.UUID
	unitMap     valueobject.UnitMap
	unitMapSize valueobject.MapSize
}

func NewGame(unitMap valueobject.UnitMap) Game {
	id, _ := uuid.NewUUID()
	unitMapSize := getUnitMapSize(unitMap)
	return Game{
		id:          id,
		unitMap:     unitMap,
		unitMapSize: unitMapSize,
	}
}

func (g *Game) GetId() uuid.UUID {
	return g.id
}

func (g *Game) GetUnitMap() valueobject.UnitMap {
	return g.unitMap
}

func (g *Game) SetUnitMap(newUnitMap valueobject.UnitMap) {
	g.unitMap = newUnitMap
	g.unitMapSize = getUnitMapSize(newUnitMap)
}

func (g *Game) GetUnitMapSize() valueobject.MapSize {
	return g.unitMapSize
}
