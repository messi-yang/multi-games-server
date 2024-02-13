package linkunitmodel

type LinkUnitRepo interface {
	Add(LinkUnit) error
	Get(LinkUnitId) (LinkUnit, error)
	Update(LinkUnit) error
	Delete(LinkUnit) error
}
