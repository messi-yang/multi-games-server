package worldaccessmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type WorldMember struct {
	id                   WorldMemberId
	worldId              globalcommonmodel.WorldId
	userId               globalcommonmodel.UserId
	role                 globalcommonmodel.WorldRole
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*WorldMember)(nil)

func NewWorldMember(
	worldId globalcommonmodel.WorldId,
	userId globalcommonmodel.UserId,
	role globalcommonmodel.WorldRole,
) WorldMember {
	newWorldRole := WorldMember{
		id:                   NewWorldMemberId(uuid.New()),
		worldId:              worldId,
		userId:               userId,
		role:                 role,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	return newWorldRole
}

func LoadWorldMember(
	id WorldMemberId,
	worldId globalcommonmodel.WorldId,
	userId globalcommonmodel.UserId,
	role globalcommonmodel.WorldRole,
	createdAt time.Time,
	updatedAt time.Time,
) WorldMember {
	return WorldMember{
		id:                   id,
		worldId:              worldId,
		userId:               userId,
		role:                 role,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (worldMember *WorldMember) PopDomainEvents() []domain.DomainEvent {
	return worldMember.domainEventCollector.PopAll()
}

func (worldMember *WorldMember) GetId() WorldMemberId {
	return worldMember.id
}

func (worldMember *WorldMember) GeWorldId() globalcommonmodel.WorldId {
	return worldMember.worldId
}

func (worldMember *WorldMember) GeUserId() globalcommonmodel.UserId {
	return worldMember.userId
}

func (worldMember *WorldMember) GetRole() globalcommonmodel.WorldRole {
	return worldMember.role
}

func (worldMember *WorldMember) GetCreatedAt() time.Time {
	return worldMember.createdAt
}

func (worldMember *WorldMember) GetUpdatedAt() time.Time {
	return worldMember.updatedAt
}
