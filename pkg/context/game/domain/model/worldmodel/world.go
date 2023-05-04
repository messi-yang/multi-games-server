package worldmodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type World struct {
	id                   commonmodel.WorldId
	gamerId              commonmodel.GamerId
	name                 string
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*World)(nil)

func NewWorld(id commonmodel.WorldId, gamerId commonmodel.GamerId, name string) World {
	return World{
		id:                   id,
		gamerId:              gamerId,
		name:                 name,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (world *World) PopDomainEvents() []domain.DomainEvent {
	return world.domainEventCollector.PopAll()
}

func (world *World) GetId() commonmodel.WorldId {
	return world.id
}

func (world *World) GetGamerId() commonmodel.GamerId {
	return world.gamerId
}

func (world *World) GetName() string {
	return world.name
}

func (world *World) ChangeName(name string) {
	world.name = name
}
