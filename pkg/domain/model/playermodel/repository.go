package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

var (
	ErrPlayerNotFound   = errors.New("player of the given id not found")
	ErrSomethinHappened = errors.New("some unexpected error happened")
)

type Repository interface {
	Add(PlayerAgg) error
	Get(PlayerIdVo) (PlayerAgg, error)
	GetPlayerAt(worldmodel.WorldIdVo, commonmodel.PositionVo) (PlayerAgg, bool, error)
	GetPlayersAround(worldmodel.WorldIdVo, commonmodel.PositionVo) ([]PlayerAgg, error)
	Update(PlayerAgg) error
	Delete(PlayerIdVo) error
}
