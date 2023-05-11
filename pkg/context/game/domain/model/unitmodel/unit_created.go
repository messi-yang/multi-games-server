package unitmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
)

type UnitCreated struct {
	occurredOn time.Time
	unitId     commonmodel.UnitId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*UnitCreated)(nil)

func NewUnitCreated(unitId commonmodel.UnitId) UnitCreated {
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

func (domainEvent UnitCreated) GetUnitId() commonmodel.UnitId {
	return domainEvent.unitId
}
