package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Game struct {
	id      GameId
	mapSize commonmodel.Size
	map_    commonmodel.Map
}

func NewGame(id GameId, map_ commonmodel.Map) Game {
	return Game{
		id:      id,
		mapSize: map_.GetSize(),
		map_:    map_,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetMap() commonmodel.Map {
	return game.map_
}

func (game *Game) GetMapSize() commonmodel.Size {
	return game.mapSize
}
