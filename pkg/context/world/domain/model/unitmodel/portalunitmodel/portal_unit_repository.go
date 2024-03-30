package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type PortalUnitRepo interface {
	Add(PortalUnit) error
	Get(PortalUnitId) (PortalUnit, error)
	Update(PortalUnit) error
	Delete(PortalUnit) error
	GetTopLeftMostUnitWithoutTarget(worldId globalcommonmodel.WorldId) (*PortalUnit, error)
}
