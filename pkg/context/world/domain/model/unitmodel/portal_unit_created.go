package unitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type PortalUnitCreated struct {
	occurredOn time.Time
	portalUnit PortalUnit
}

// Interface Implementation Check
var _ domain.DomainEvent = (*PortalUnitCreated)(nil)

func NewPortalUnitCreated(
	portalUnit PortalUnit,
) PortalUnitCreated {
	return PortalUnitCreated{
		occurredOn: time.Now(),
		portalUnit: portalUnit,
	}
}

func (domainEvent PortalUnitCreated) GetEventName() string {
	return "PORTAL_UNIT_CREATED"
}

func (domainEvent PortalUnitCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent PortalUnitCreated) GetPortalUnit() PortalUnit {
	return domainEvent.portalUnit
}
