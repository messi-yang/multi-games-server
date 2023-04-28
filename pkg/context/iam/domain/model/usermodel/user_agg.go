package usermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type UserAgg struct {
	id           sharedkernelmodel.UserIdVo
	emailAddress string
	username     string
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*UserAgg)(nil)

func NewUserAgg(
	id sharedkernelmodel.UserIdVo,
	emailAddress string,
	username string,
) UserAgg {
	newUser := UserAgg{id: id, emailAddress: emailAddress, username: username, domainEvents: []domainmodel.DomainEvent{}}
	newUser.AddDomainEvent(NewUserCreatedDomainEvent(id))
	return newUser
}

func (agg *UserAgg) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	agg.domainEvents = append(agg.domainEvents, domainEvent)
}

func (agg *UserAgg) GetDomainEvents() []domainmodel.DomainEvent {
	return agg.domainEvents
}

func (agg *UserAgg) GetId() sharedkernelmodel.UserIdVo {
	return agg.id
}

func (agg *UserAgg) GetEmailAddress() string {
	return agg.emailAddress
}

func (agg *UserAgg) GetUsername() string {
	return agg.username
}
