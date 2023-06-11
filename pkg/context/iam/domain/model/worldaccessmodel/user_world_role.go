package worldaccessmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type UserWorldRole struct {
	id                   UserWorldRoleId
	worldId              sharedkernelmodel.WorldId
	userId               sharedkernelmodel.UserId
	worldRole            WorldRole
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*UserWorldRole)(nil)

func NewUserWorldRole(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	worldRole WorldRole,
) UserWorldRole {
	newWorldRole := UserWorldRole{
		id:                   NewUserWorldRoleId(uuid.New()),
		worldId:              worldId,
		userId:               userId,
		worldRole:            worldRole,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	return newWorldRole
}

func LoadWorldRole(
	id UserWorldRoleId,
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	worldRole WorldRole,
	createdAt time.Time,
	updatedAt time.Time,
) UserWorldRole {
	return UserWorldRole{
		id:                   id,
		worldId:              worldId,
		userId:               userId,
		worldRole:            worldRole,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (userWorldRole *UserWorldRole) PopDomainEvents() []domain.DomainEvent {
	return userWorldRole.domainEventCollector.PopAll()
}

func (userWorldRole *UserWorldRole) GetId() UserWorldRoleId {
	return userWorldRole.id
}

func (userWorldRole *UserWorldRole) GeWorldId() sharedkernelmodel.WorldId {
	return userWorldRole.worldId
}

func (userWorldRole *UserWorldRole) GeUserId() sharedkernelmodel.UserId {
	return userWorldRole.userId
}

func (userWorldRole *UserWorldRole) GetWorldRole() WorldRole {
	return userWorldRole.worldRole
}

func (userWorldRole *UserWorldRole) GetCreatedAt() time.Time {
	return userWorldRole.createdAt
}

func (userWorldRole *UserWorldRole) GetUpdatedAt() time.Time {
	return userWorldRole.updatedAt
}
