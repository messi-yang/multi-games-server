package itemmodel

type Repo interface {
	GetAll() []ItemAgg
	Get(itemId ItemIdVo) (ItemAgg, error)
}
