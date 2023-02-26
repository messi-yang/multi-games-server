package gamemodel

import (
	"errors"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type GameAgg struct {
	id GameIdVo
}

func NewGameAgg(id GameIdVo) GameAgg {
	return GameAgg{
		id: id,
	}
}

func (game *GameAgg) GetId() GameIdVo {
	return game.id
}
