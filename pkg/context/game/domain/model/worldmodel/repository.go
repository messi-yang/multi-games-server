package worldmodel

type Repository interface {
	Add(WorldAgg) error
	Get(WorldIdVo) (WorldAgg, error)
	GetAll() ([]WorldAgg, error)

	ReadLockAccess(WorldIdVo) (rUnlocker func())
	LockAccess(WorldIdVo) (unlocker func())
}
