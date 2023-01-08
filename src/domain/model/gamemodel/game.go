package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Game struct {
	id      GameId
	mapSize commonmodel.MapSize
	mapVo   commonmodel.Map
}

func NewGame(id GameId, mapVo commonmodel.Map) Game {
	return Game{
		id:      id,
		mapSize: mapVo.GetMapSize(),
		mapVo:   mapVo,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetMap() commonmodel.Map {
	return game.mapVo
}

func (game *Game) GetMapSize() commonmodel.MapSize {
	return game.mapSize
}
