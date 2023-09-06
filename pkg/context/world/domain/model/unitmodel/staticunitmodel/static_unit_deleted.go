package staticunitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type StaticUnitDeleted struct {
	occurredOn time.Time
	portalUnit StaticUnit
}

// Interface Implementation Check
var _ domain.DomainEvent = (*StaticUnitDeleted)(nil)

func NewStaticUnitDeleted(
	portalUnit StaticUnit,
) StaticUnitDeleted {
	return StaticUnitDeleted{
		occurredOn: time.Now(),
		portalUnit: portalUnit,
	}
}

func (domainEvent StaticUnitDeleted) GetEventName() string {
	return "STATIC_UNIT_DELETED"
}

func (domainEvent StaticUnitDeleted) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent StaticUnitDeleted) GetStaticUnit() StaticUnit {
	return domainEvent.portalUnit
}
