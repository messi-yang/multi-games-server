package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
)

type PortalUnitRepo interface {
	Add(PortalUnit) error
	Get(unitmodel.UnitId) (PortalUnit, error)
	Update(PortalUnit) error
	Delete(PortalUnit) error
	GetTopLeftMostUnitWithoutTarget(worldId globalcommonmodel.WorldId) (*PortalUnit, error)
}
