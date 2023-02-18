package gamemodel

import (
	"errors"
)

var (
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrLocationHasPlayer             = errors.New("the location has player")
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
