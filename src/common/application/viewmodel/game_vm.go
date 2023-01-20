package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
)

type GameVm struct {
	Id      string `json:"id"`
	MapSize SizeVm `json:"mapSize"`
}

func NewGameVm(game gamemodel.GameAgg) GameVm {
	return GameVm{
		Id:      game.GetId().ToString(),
		MapSize: NewSizeVm(game.GetMapSize()),
	}
}
