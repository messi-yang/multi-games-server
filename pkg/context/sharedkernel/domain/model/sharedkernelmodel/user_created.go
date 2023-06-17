package sharedkernelmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type UserCreated struct {
	occurredOn time.Time
	userId     UserId
}

// Interface Implementation Check
var _ domain.DomainEvent = (*UserCreated)(nil)

func NewUserCreated(userId UserId) UserCreated {
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

func (domainEvent UserCreated) GetUserId() UserId {
	return domainEvent.userId
}
