package usermodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type UserCreated struct {
	occurredOn time.Time
	userId     sharedkernelmodel.UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*UserCreated)(nil)

func NewUserCreated(userId sharedkernelmodel.UserId) UserCreated {
	return UserCreated{
		occurredOn: time.Now(),
		userId:     userId,
	}
}

func (domainEvent UserCreated) GetEventName() string {
	return "USER_CREATED"
}

func (domainEvent UserCreated) GetOccurredOn() time.Time {
	return domainEvent.occurredOn
}

func (domainEvent UserCreated) GetUserId() sharedkernelmodel.UserId {
	return domainEvent.userId
}
