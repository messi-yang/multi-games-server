package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type GameAgg struct {
	id      GameIdVo
	mapSize commonmodel.SizeVo
	map_    MapVo
}

func NewGameAgg(id GameIdVo, map_ MapVo) GameAgg {
	return GameAgg{
		id:      id,
		mapSize: map_.GetSize(),
		map_:    map_,
	}
}

func (game *GameAgg) GetId() GameIdVo {
	return game.id
}

func (game *GameAgg) GetMap() MapVo {
	return game.map_
}

func (game *GameAgg) GetMapSize() commonmodel.SizeVo {
	return game.mapSize
}
