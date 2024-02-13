package staticunitmodel

type StaticUnitRepo interface {
	Add(StaticUnit) error
	Update(StaticUnit) error
	Get(StaticUnitId) (StaticUnit, error)
	Delete(StaticUnit) error
}
