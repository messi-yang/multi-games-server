package worldmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/google/uuid"
)

type World struct {
	id                   sharedkernelmodel.WorldId
	userId               sharedkernelmodel.UserId
	name                 string
	bound                commonmodel.Bound
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*World)(nil)

func NewWorld(
	userId sharedkernelmodel.UserId,
	name string,
	bound commonmodel.Bound,
) World {
	newWorld := World{
		id:                   sharedkernelmodel.NewWorldId(uuid.New()),
		userId:               userId,
		name:                 name,
		bound:                bound,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	newWorld.domainEventCollector.Add(sharedkernelmodel.NewWorldCreated(
		newWorld.id,
		newWorld.userId,
	))
	return newWorld
}

func LoadWorld(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	name string,
	bound commonmodel.Bound,
	createdAt time.Time,
	updatedAt time.Time,
) World {
	return World{
		id:                   worldId,
		userId:               userId,
		name:                 name,
		bound:                bound,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (world *World) PopDomainEvents() []domain.DomainEvent {
	return world.domainEventCollector.PopAll()
}

func (world *World) GetId() sharedkernelmodel.WorldId {
	return world.id
}

func (world *World) GetUserId() sharedkernelmodel.UserId {
	return world.userId
}

func (world *World) GetName() string {
	return world.name
}

func (world *World) GetBound() commonmodel.Bound {
	return world.bound
}

func (world *World) ChangeName(name string) {
	world.name = name
}

func (world *World) GetCreatedAt() time.Time {
	return world.createdAt
}

func (world *World) GetUpdatedAt() time.Time {
	return world.updatedAt
}
