package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
)

var (
	ErrPlayerNotFound = errors.New("player with id not found")
)

type Repo interface {
	Add(PlayerAgg) error
	Get(PlayerIdVo) (PlayerAgg, error)
	GetPlayerAt(commonmodel.LocationVo) (PlayerAgg, error)
	Update(PlayerAgg) error
	GetAll() []PlayerAgg
	Delete(PlayerIdVo)

	ReadLockAccess(PlayerIdVo) (rUnlocker func())
	LockAccess(PlayerIdVo) (unlocker func())
}
