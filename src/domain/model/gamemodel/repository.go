package gamemodel

type GameRepo interface {
	Add(GameAgr) error
	Get(GameIdVo) (GameAgr, error)
	Update(GameIdVo, GameAgr) error
	GetAll() ([]GameAgr, error)

	ReadLockAccess(GameIdVo) (rUnlocker func(), err error)
	LockAccess(GameIdVo) (unlocker func(), err error)
}
