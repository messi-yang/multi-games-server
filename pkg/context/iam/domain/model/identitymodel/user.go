package identitymodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type User struct {
	id                   sharedkernelmodel.UserId
	emailAddress         string
	username             string
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*User)(nil)

func NewUser(
	id sharedkernelmodel.UserId,
	emailAddress string,
	username string,
) User {
	newUser := User{id: id, emailAddress: emailAddress, username: username, domainEventCollector: domain.NewDomainEventCollector()}
	newUser.domainEventCollector.Add(sharedkernelmodel.NewUserCreated(id))
	return newUser
}

func (user *User) PopDomainEvents() []domain.DomainEvent {
	return user.domainEventCollector.PopAll()
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
