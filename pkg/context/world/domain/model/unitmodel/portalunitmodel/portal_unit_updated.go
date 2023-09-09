package portalunitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type PortalUnitUpdated struct {
	occurredOn time.Time
	portalUnit PortalUnit
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PortalUnitUpdated)(nil)

func NewPortalUnitUpdated(
	portalUnit PortalUnit,
) PortalUnitUpdated {
	return PortalUnitUpdated{
		occurredOn: time.Now(),
		portalUnit: portalUnit,
	}
}

func (domainEvent PortalUnitUpdated) GetEventName() string {
	return "PORTAL_UNIT_UPDATED"
}

func (domainEvent PortalUnitUpdated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent PortalUnitUpdated) GetPortalUnit() PortalUnit {
	return domainEvent.portalUnit
}
