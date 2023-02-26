package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
)

var (
	ErrPlayerNotFound = errors.New("player with id not found")
)

type Repo interface {
	Add(PlayerAgg) error
	Get(PlayerIdVo) (PlayerAgg, error)
	GetPlayerAt(gamemodel.GameIdVo, commonmodel.PositionVo) (PlayerAgg, bool, error)
	GetPlayersAround(gamemodel.GameIdVo, commonmodel.PositionVo) ([]PlayerAgg, error)
	Update(PlayerAgg) error
	Delete(PlayerIdVo) error
}
