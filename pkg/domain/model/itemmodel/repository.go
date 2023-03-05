package itemmodel

type Repo interface {
	GetAll() ([]ItemAgg, error)
	Get(itemId ItemIdVo) (ItemAgg, error)
}
