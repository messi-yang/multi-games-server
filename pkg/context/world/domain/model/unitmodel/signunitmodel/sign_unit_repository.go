package signunitmodel

type SignUnitRepo interface {
	Add(SignUnit) error
	Get(SignUnitId) (SignUnit, error)
	Update(SignUnit) error
	Delete(SignUnit) error
}
