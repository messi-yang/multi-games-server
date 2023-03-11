package itemmodel

type Repository interface {
	GetAll() ([]ItemAgg, error)
	Get(itemId ItemIdVo) (ItemAgg, error)
}
