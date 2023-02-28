package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

var (
	ErrPlayerNotFound = errors.New("player with id not found")
)

type Repo interface {
	Add(PlayerAgg) error
	Get(PlayerIdVo) (PlayerAgg, error)
	GetPlayerAt(worldmodel.WorldIdVo, commonmodel.PositionVo) (PlayerAgg, bool, error)
	GetPlayersAround(worldmodel.WorldIdVo, commonmodel.PositionVo) ([]PlayerAgg, error)
	Update(PlayerAgg) error
	Delete(PlayerIdVo) error
}
