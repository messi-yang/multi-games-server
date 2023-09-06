package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PortalUnitRepo interface {
	Add(PortalUnit) error
	Get(unitmodel.UnitId) (PortalUnit, error)
	Update(PortalUnit) error
	Delete(PortalUnit) error
	GetRandomPortalUnit(worldId globalcommonmodel.WorldId) (*PortalUnit, error)
	FindPortalUnitWithTargetPosition(worldId globalcommonmodel.WorldId, targetPosition worldcommonmodel.Position) (*PortalUnit, error)
}
