package livegamemodel

type Repo interface {
	Add(LiveGameAgg) error
	Get(LiveGameIdVo) (LiveGameAgg, error)
	Update(LiveGameIdVo, LiveGameAgg) error
	GetAll() []LiveGameAgg

	ReadLockAccess(LiveGameIdVo) (rUnlocker func())
	LockAccess(LiveGameIdVo) (unlocker func())
}
