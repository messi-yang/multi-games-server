package gamemodel

type GameRepo interface {
	Add(Game) error
	Get(GameId) (Game, error)
	Update(GameId, Game) error
	GetAll() ([]Game, error)

	ReadLockAccess(GameId) (rUnlocker func(), err error)
	LockAccess(GameId) (unlocker func(), err error)
}
