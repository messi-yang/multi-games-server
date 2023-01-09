package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Game struct {
	id        GameId
	dimension commonmodel.Dimension
	mapVo     commonmodel.Map
}

func NewGame(id GameId, mapVo commonmodel.Map) Game {
	return Game{
		id:        id,
		dimension: mapVo.GetDimension(),
		mapVo:     mapVo,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetMap() commonmodel.Map {
	return game.mapVo
}

func (game *Game) GetDimension() commonmodel.Dimension {
	return game.dimension
}
