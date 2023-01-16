package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type GameAgr struct {
	id      GameIdVo
	mapSize commonmodel.SizeVo
	map_    MapVo
}

func NewGameAgr(id GameIdVo, map_ MapVo) GameAgr {
	return GameAgr{
		id:      id,
		mapSize: map_.GetSize(),
		map_:    map_,
	}
}

func (game *GameAgr) GetId() GameIdVo {
	return game.id
}

func (game *GameAgr) GetMap() MapVo {
	return game.map_
}

func (game *GameAgr) GetMapSize() commonmodel.SizeVo {
	return game.mapSize
}
