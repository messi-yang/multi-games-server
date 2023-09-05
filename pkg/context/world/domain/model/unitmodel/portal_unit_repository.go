package unitmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type PortalUnitRepo interface {
	Add(PortalUnit) error
	Get(UnitId) (PortalUnit, error)
	Delete(PortalUnit) error
	GetRandomPortalUnit(worldId globalcommonmodel.WorldId) (*PortalUnit, error)
}
