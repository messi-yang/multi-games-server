package worldmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type Repo interface {
	Add(World) error
	Get(commonmodel.WorldId) (World, error)
	GetAll() ([]World, error)
	ReadLockAccess(commonmodel.WorldId) (rUnlocker func())
	LockAccess(commonmodel.WorldId) (unlocker func())
}
