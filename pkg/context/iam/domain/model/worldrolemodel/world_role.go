package worldrolemodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type WorldRole struct {
	id                   WorldRoleId
	worldId              sharedkernelmodel.WorldId
	userId               sharedkernelmodel.UserId
	name                 WorldRoleName
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*WorldRole)(nil)

func NewWorldRole(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	name WorldRoleName,
) WorldRole {
	newWorldRole := WorldRole{
		id:                   NewWorldRoleId(uuid.New()),
		worldId:              worldId,
		userId:               userId,
		name:                 name,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	return newWorldRole
}

func LoadWorldRole(
	id WorldRoleId,
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	name WorldRoleName,
	createdAt time.Time,
	updatedAt time.Time,
) WorldRole {
	return WorldRole{
		id:                   id,
		worldId:              worldId,
		userId:               userId,
		name:                 name,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (worldRole *WorldRole) PopDomainEvents() []domain.DomainEvent {
	return worldRole.domainEventCollector.PopAll()
}

func (worldRole *WorldRole) GetId() WorldRoleId {
	return worldRole.id
}

func (worldRole *WorldRole) GeWorldId() sharedkernelmodel.WorldId {
	return worldRole.worldId
}

func (worldRole *WorldRole) GeUserId() sharedkernelmodel.UserId {
	return worldRole.userId
}

func (worldRole *WorldRole) GetName() WorldRoleName {
	return worldRole.name
}

func (worldRole *WorldRole) GetCreatedAt() time.Time {
	return worldRole.createdAt
}

func (worldRole *WorldRole) GetUpdatedAt() time.Time {
	return worldRole.updatedAt
}
