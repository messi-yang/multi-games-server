package worldmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/domainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
)

type World struct {
	id                   globalcommonmodel.WorldId
	userId               globalcommonmodel.UserId
	name                 string
	bound                worldcommonmodel.Bound
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.DomainEventDispatchableAggregate = (*World)(nil)

func NewWorld(
	userId globalcommonmodel.UserId,
	name string,
	bound worldcommonmodel.Bound,
) World {
	newWorld := World{
		id:                   globalcommonmodel.NewWorldId(uuid.New()),
		userId:               userId,
		name:                 name,
		bound:                bound,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
	newWorld.domainEventCollector.Add(domainevent.NewWorldCreated(
		newWorld.id,
		newWorld.userId,
	))
	return newWorld
}

func LoadWorld(
	worldId globalcommonmodel.WorldId,
	userId globalcommonmodel.UserId,
	name string,
	bound worldcommonmodel.Bound,
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

func (world *World) GetId() globalcommonmodel.WorldId {
	return world.id
}

func (world *World) GetUserId() globalcommonmodel.UserId {
	return world.userId
}

func (world *World) GetName() string {
	return world.name
}

func (world *World) GetBound() worldcommonmodel.Bound {
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

func (world *World) Delete() {
	world.domainEventCollector.Add(domainevent.NewWorldCreated(
		world.id,
		world.userId,
	))
}
