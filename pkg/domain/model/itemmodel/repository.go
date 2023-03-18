package itemmodel

type Repository interface {
	GetAll() ([]ItemAgg, error)
	Get(itemId ItemIdVo) (ItemAgg, error)
	Add(item ItemAgg) error
}
