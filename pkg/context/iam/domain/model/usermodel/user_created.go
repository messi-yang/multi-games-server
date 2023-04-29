package usermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type UserCreated struct {
	OccurredOn time.Time
	UserId     sharedkernelmodel.UserId
}

// Interface Implementation Check
var _ domainmodel.DomainEvent = (*UserCreated)(nil)

func NewUserCreated(userId sharedkernelmodel.UserId) UserCreated {
	return UserCreated{
		OccurredOn: time.Now(),
		UserId:     userId,
	}
}

func (domainEvent UserCreated) GetName() string {
	return "USER_CREATED"
}

func (domainEvent UserCreated) GetOccurredOn() time.Time {
	return domainEvent.OccurredOn
}
