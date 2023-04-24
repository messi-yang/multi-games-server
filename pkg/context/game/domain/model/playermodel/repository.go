package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

var (
	ErrPlayerNotFound   = errors.New("player of the given id not found")
	ErrSomethinHappened = errors.New("some unexpected error happened")
)

type Repository interface {
	Add(PlayerAgg) error
	Get(commonmodel.PlayerIdVo) (PlayerAgg, error)
	FindPlayerAt(commonmodel.WorldIdVo, commonmodel.PositionVo) (player PlayerAgg, found bool, err error)
	GetPlayersAround(commonmodel.WorldIdVo, commonmodel.PositionVo) ([]PlayerAgg, error)
	Update(PlayerAgg) error
	Delete(commonmodel.PlayerIdVo) error
}
