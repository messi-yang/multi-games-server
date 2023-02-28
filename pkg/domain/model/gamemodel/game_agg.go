package gamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type GameAgg struct {
	id     GameIdVo
	userId usermodel.UserIdVo
	name   string
}

func NewGameAgg(id GameIdVo, userId usermodel.UserIdVo) GameAgg {
	return GameAgg{
		id:     id,
		userId: userId,
		name:   "Hello World",
	}
}

func (game *GameAgg) GetId() GameIdVo {
	return game.id
}

func (game *GameAgg) GetUserId() usermodel.UserIdVo {
	return game.userId
}

func (game *GameAgg) GetName() string {
	return game.name
}
