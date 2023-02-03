package itemmodel

type Repo interface {
	GetAllItems() []ItemAgg
	Get(itemId ItemIdVo) (ItemAgg, error)
}
