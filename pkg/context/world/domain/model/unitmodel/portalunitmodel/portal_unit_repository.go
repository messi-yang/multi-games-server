package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PortalUnitRepo interface {
	Add(PortalUnit) error
	Get(PortalUnitId) (PortalUnit, error)
	Find(globalcommonmodel.WorldId, worldcommonmodel.Position) (*PortalUnit, error)
	Update(PortalUnit) error
	Delete(PortalUnit) error
	GetTopLeftMostUnitWithoutTarget(worldId globalcommonmodel.WorldId) (*PortalUnit, error)
}
