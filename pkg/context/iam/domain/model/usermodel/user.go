package usermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type User struct {
	id           sharedkernelmodel.UserId
	emailAddress string
	username     string
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*User)(nil)

func NewUser(
	id sharedkernelmodel.UserId,
	emailAddress string,
	username string,
) User {
	newUser := User{id: id, emailAddress: emailAddress, username: username, domainEvents: []domainmodel.DomainEvent{}}
	newUser.AddDomainEvent(NewUserCreated(id))
	return newUser
}

func (user *User) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	user.domainEvents = append(user.domainEvents, domainEvent)
}

func (user *User) GetDomainEvents() []domainmodel.DomainEvent {
	return user.domainEvents
}

func (user *User) GetId() sharedkernelmodel.UserId {
	return user.id
}

func (user *User) GetEmailAddress() string {
	return user.emailAddress
}

func (user *User) GetUsername() string {
	return user.username
}
