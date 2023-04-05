package worldmodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type WorldAgg struct {
	id     WorldIdVo
	userId usermodel.UserIdVo
	name   string
}

func NewWorldAgg(id WorldIdVo, userId usermodel.UserIdVo) WorldAgg {
	return WorldAgg{
		id:     id,
		userId: userId,
		name:   "Hello World",
	}
}

func (agg *WorldAgg) GetId() WorldIdVo {
	return agg.id
}

func (agg *WorldAgg) GetUserId() usermodel.UserIdVo {
	return agg.userId
}

func (agg *WorldAgg) GetName() string {
	return agg.name
}
