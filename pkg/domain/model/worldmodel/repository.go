package worldmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"

type Repository interface {
	Add(WorldAgg) error
	Get(WorldIdVo) (WorldAgg, error)
	GetWorldOfUser(usermodel.UserIdVo) (WorldAgg, bool, error)
	GetAll() ([]WorldAgg, error)

	ReadLockAccess(WorldIdVo) (rUnlocker func())
	LockAccess(WorldIdVo) (unlocker func())
}
