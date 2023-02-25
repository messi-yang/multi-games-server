package gamemodel

type Repo interface {
	Add(GameAgg) error
	Get(GameIdVo) (GameAgg, bool, error)
	Update(GameAgg) error
	GetAll() ([]GameAgg, error)

	ReadLockAccess(GameIdVo) (rUnlocker func())
	LockAccess(GameIdVo) (unlocker func())
}
