package playermodel

type Repo interface {
	Add(PlayerAgg)
	GetAll() []PlayerAgg
	Remove(PlayerIdVo)
}
