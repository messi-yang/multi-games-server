package staticunitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type StaticUnitUpdated struct {
	occurredOn time.Time
	portalUnit StaticUnit
}

// Interface Implementation Check
var _ domain.DomainEvent = (*StaticUnitUpdated)(nil)

func NewStaticUnitUpdated(
	portalUnit StaticUnit,
) StaticUnitUpdated {
	return StaticUnitUpdated{
		occurredOn: time.Now(),
		portalUnit: portalUnit,
	}
}

func (domainEvent StaticUnitUpdated) GetEventName() string {
	return "STATIC_UNIT_ROTATED"
}

func (domainEvent StaticUnitUpdated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent StaticUnitUpdated) GetStaticUnit() StaticUnit {
	return domainEvent.portalUnit
}
