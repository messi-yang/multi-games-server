package livegamemodel

type LiveGameRepository interface {
	Add(LiveGame) error
	Get(LiveGameId) (LiveGame, error)
	Update(LiveGameId, LiveGame) error
	GetAll() []LiveGame

	ReadLockAccess(LiveGameId) (rUnlocker func(), err error)
	LockAccess(LiveGameId) (unlocker func(), err error)
}
