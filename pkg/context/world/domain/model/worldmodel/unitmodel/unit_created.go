package unitmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type UnitCreated struct {
	occurredOn time.Time
	unitId     UnitId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*UnitCreated)(nil)

func NewUnitCreated(unitId UnitId) UnitCreated {
	return UnitCreated{
		occurredOn: time.Now(),
		unitId:     unitId,
	}
}

func (domainEvent UnitCreated) GetEventName() string {
	return "UNIT_CREATED"
}

func (domainEvent UnitCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent UnitCreated) GetUnitId() UnitId {
	return domainEvent.unitId
}
