package colorunitmodel

type ColorUnitRepo interface {
	Add(ColorUnit) error
	Get(ColorUnitId) (ColorUnit, error)
	Update(ColorUnit) error
	Delete(ColorUnit) error
}
