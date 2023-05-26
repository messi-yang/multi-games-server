package worldmodel

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type World struct {
	id                   commonmodel.WorldId
	userId               sharedkernelmodel.UserId
	name                 string
	createdAt            time.Time
	updatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*World)(nil)

func NewWorld(userId sharedkernelmodel.UserId, name string) World {
	return World{
		id:                   commonmodel.NewWorldId(uuid.New()),
		userId:               userId,
		name:                 name,
		createdAt:            time.Now(),
		updatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func LoadWorld(
	worldId commonmodel.WorldId,
	userId sharedkernelmodel.UserId,
	name string,
	createdAt time.Time,
	updatedAt time.Time,
) World {
	return World{
		id:                   worldId,
		userId:               userId,
		name:                 name,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (world *World) PopDomainEvents() []domain.DomainEvent {
	return world.domainEventCollector.PopAll()
}

func (world *World) GetId() commonmodel.WorldId {
	return world.id
}

func (world *World) GetUserId() sharedkernelmodel.UserId {
	return world.userId
}

func (world *World) GetName() string {
	return world.name
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
