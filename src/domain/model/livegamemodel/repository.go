package livegamemodel

type Repo interface {
	Add(LiveGameAgr) error
	Get(LiveGameIdVo) (LiveGameAgr, error)
	Update(LiveGameIdVo, LiveGameAgr) error
	GetAll() []LiveGameAgr

	ReadLockAccess(LiveGameIdVo) (rUnlocker func())
	LockAccess(LiveGameIdVo) (unlocker func())
}
