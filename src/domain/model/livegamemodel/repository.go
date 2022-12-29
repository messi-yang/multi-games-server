package livegamemodel

type Repo interface {
	Add(LiveGame) error
	Get(LiveGameId) (LiveGame, error)
	Update(LiveGameId, LiveGame) error
	GetAll() []LiveGame

	ReadLockAccess(LiveGameId) (rUnlocker func())
	LockAccess(LiveGameId) (unlocker func())
}
