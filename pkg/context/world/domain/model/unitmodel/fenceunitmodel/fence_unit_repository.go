package fenceunitmodel

type FenceUnitRepo interface {
	Add(FenceUnit) error
	Update(FenceUnit) error
	Get(FenceUnitId) (FenceUnit, error)
	Delete(FenceUnit) error
}
