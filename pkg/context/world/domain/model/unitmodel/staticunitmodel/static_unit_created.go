package staticunitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type StaticUnitCreated struct {
	occurredOn time.Time
	portalUnit StaticUnit
}

// Interface Implementation Check
var _ domain.DomainEvent = (*StaticUnitCreated)(nil)

func NewStaticUnitCreated(
	portalUnit StaticUnit,
) StaticUnitCreated {
	return StaticUnitCreated{
		occurredOn: time.Now(),
		portalUnit: portalUnit,
	}
}

func (domainEvent StaticUnitCreated) GetEventName() string {
	return "STATIC_UNIT_CREATED"
}

func (domainEvent StaticUnitCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent StaticUnitCreated) GetStaticUnit() StaticUnit {
	return domainEvent.portalUnit
}
