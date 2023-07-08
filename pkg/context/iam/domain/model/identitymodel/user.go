package identitymodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type User struct {
	id                   sharedkernelmodel.UserId
	emailAddress         sharedkernelmodel.EmailAddress
	username             sharedkernelmodel.Username
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*User)(nil)

func NewUser(
	id sharedkernelmodel.UserId,
	emailAddress sharedkernelmodel.EmailAddress,
	username sharedkernelmodel.Username,
) User {
	newUser := User{
		id:                   id,
		emailAddress:         emailAddress,
		username:             username,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	newUser.domainEventCollector.Add(sharedkernelmodel.NewUserCreated(id))
	return newUser
}

func LoadUser(
	id sharedkernelmodel.UserId,
	emailAddress sharedkernelmodel.EmailAddress,
	username sharedkernelmodel.Username,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	newUser := User{
		id:                   id,
		emailAddress:         emailAddress,
		username:             username,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	return newUser
}

func (user *User) PopDomainEvents() []domain.DomainEvent {
	return user.domainEventCollector.PopAll()
}

func (user *User) GetId() sharedkernelmodel.UserId {
	return user.id
}

func (user *User) GetEmailAddress() sharedkernelmodel.EmailAddress {
	return user.emailAddress
}

func (user *User) GetUsername() sharedkernelmodel.Username {
	return user.username
}

func (user *User) GetCreatedAt() time.Time {
	return user.createdAt
}

func (user *User) GetUpdatedAt() time.Time {
	return user.updatedAt
}
