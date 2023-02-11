package gamemodel

type Repo interface {
	Add(GameAgg) error
	Get(GameIdVo) (GameAgg, error)
	Update(GameIdVo, GameAgg) error
	GetAll() []GameAgg

	ReadLockAccess(GameIdVo) (rUnlocker func())
	LockAccess(GameIdVo) (unlocker func())
}
