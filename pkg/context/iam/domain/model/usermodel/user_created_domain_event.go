package usermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type UserCreatedDomainEvent struct {
	OccurredOn time.Time
	UserId     sharedkernelmodel.UserId
}

// Interface Implementation Check
var _ domainmodel.DomainEvent = (*UserCreatedDomainEvent)(nil)

func NewUserCreatedDomainEvent(userId sharedkernelmodel.UserId) UserCreatedDomainEvent {
	return UserCreatedDomainEvent{
		OccurredOn: time.Now(),
		UserId:     userId,
	}
}

func (domainEvent UserCreatedDomainEvent) GetOccurredOn() time.Time {
	return domainEvent.OccurredOn
}
