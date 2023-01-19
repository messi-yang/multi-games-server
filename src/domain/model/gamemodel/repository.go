package gamemodel

type GameRepo interface {
	Add(GameAgg) error
	Get(GameIdVo) (GameAgg, error)
	Update(GameIdVo, GameAgg) error
	GetAll() ([]GameAgg, error)

	ReadLockAccess(GameIdVo) (rUnlocker func(), err error)
	LockAccess(GameIdVo) (unlocker func(), err error)
}
