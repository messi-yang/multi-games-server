package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Game struct {
	id      GameId
	mapSize commonmodel.MapSize
	unitMap commonmodel.UnitMap
}

func NewGame(id GameId, unitMap commonmodel.UnitMap) Game {
	return Game{
		id:      id,
		mapSize: unitMap.GetMapSize(),
		unitMap: unitMap,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetUnitMap() commonmodel.UnitMap {
	return game.unitMap
}

func (game *Game) GetMapSize() commonmodel.MapSize {
	return game.mapSize
}
