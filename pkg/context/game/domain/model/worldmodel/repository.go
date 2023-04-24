package worldmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type Repo interface {
	Add(WorldAgg) error
	Get(commonmodel.WorldIdVo) (WorldAgg, error)
	GetAll() ([]WorldAgg, error)

	ReadLockAccess(commonmodel.WorldIdVo) (rUnlocker func())
	LockAccess(commonmodel.WorldIdVo) (unlocker func())
}
