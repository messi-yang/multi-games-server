package unitmodel

type PortalUnitRepo interface {
	Add(PortalUnit) error
	Get(PortalUnitId) (PortalUnit, error)
	Delete(PortalUnit) error
}
