package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Game struct {
	id             GameId
	unitMapMapSize commonmodel.MapSize
	unitMap        commonmodel.UnitMap
}

func NewGame(id GameId, unitMap commonmodel.UnitMap) Game {
	return Game{
		id:             id,
		unitMapMapSize: unitMap.GetMapSize(),
		unitMap:        unitMap,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetUnitMap() commonmodel.UnitMap {
	return game.unitMap
}

func (game *Game) GetUnitMapMapSize() commonmodel.MapSize {
	return game.unitMapMapSize
}
