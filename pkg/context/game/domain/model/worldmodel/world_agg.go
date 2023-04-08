package worldmodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type WorldAgg struct {
	id      WorldIdVo
	gamerId gamermodel.GamerIdVo
	name    string
}

func NewWorldAgg(id WorldIdVo, gamerId gamermodel.GamerIdVo) WorldAgg {
	return WorldAgg{
		id:      id,
		gamerId: gamerId,
		name:    "Hello World",
	}
}

func (agg *WorldAgg) GetId() WorldIdVo {
	return agg.id
}

func (agg *WorldAgg) GetGamerId() gamermodel.GamerIdVo {
	return agg.gamerId
}

func (agg *WorldAgg) GetName() string {
	return agg.name
}
