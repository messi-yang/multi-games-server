package unitmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
)

type UnitDeleted struct {
	occurredOn time.Time
	unitId     commonmodel.UnitId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*UnitDeleted)(nil)

func NewUnitDeleted(unitId commonmodel.UnitId) UnitDeleted {
	return UnitDeleted{
		occurredOn: time.Now(),
		unitId:     unitId,
	}
}

func (domainEvent UnitDeleted) GetEventName() string {
	return "UNIT_DELETED"
}

func (domainEvent UnitDeleted) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent UnitDeleted) GetUnitId() commonmodel.UnitId {
	return domainEvent.unitId
}
