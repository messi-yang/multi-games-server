package portalunitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type PortalUnitDeleted struct {
	occurredOn time.Time
	portalUnit PortalUnit
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PortalUnitDeleted)(nil)

func NewPortalUnitDeleted(
	portalUnit PortalUnit,
) PortalUnitDeleted {
	return PortalUnitDeleted{
		occurredOn: time.Now(),
		portalUnit: portalUnit,
	}
}

func (domainEvent PortalUnitDeleted) GetEventName() string {
	return "PORTAL_UNIT_DELETED"
}

func (domainEvent PortalUnitDeleted) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent PortalUnitDeleted) GetPortalUnit() PortalUnit {
	return domainEvent.portalUnit
}
