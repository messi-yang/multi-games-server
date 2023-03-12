package worldmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"

type Repository interface {
	Add(WorldAgg) error
	ExistsWithUserId(usermodel.UserIdVo) (bool, error)
	GetAll() ([]WorldAgg, error)

	ReadLockAccess(WorldIdVo) (rUnlocker func())
	LockAccess(WorldIdVo) (unlocker func())
}
