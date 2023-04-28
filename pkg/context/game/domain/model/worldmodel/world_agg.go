package worldmodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

var (
	ErrSomePositionsNotIncludedInMap = errors.New("some positions are not included in the unit map")
	ErrPositionHasPlayer             = errors.New("the position has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type WorldAgg struct {
	id           commonmodel.WorldIdVo
	gamerId      commonmodel.GamerIdVo
	name         string
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*WorldAgg)(nil)

func NewWorldAgg(id commonmodel.WorldIdVo, gamerId commonmodel.GamerIdVo) WorldAgg {
	return WorldAgg{
		id:           id,
		gamerId:      gamerId,
		name:         "Hello World",
		domainEvents: []domainmodel.DomainEvent{},
	}
}

func (agg *WorldAgg) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	agg.domainEvents = append(agg.domainEvents, domainEvent)
}

func (agg *WorldAgg) GetDomainEvents() []domainmodel.DomainEvent {
	return agg.domainEvents
}

func (agg *WorldAgg) GetId() commonmodel.WorldIdVo {
	return agg.id
}

func (agg *WorldAgg) GetGamerId() commonmodel.GamerIdVo {
	return agg.gamerId
}

func (agg *WorldAgg) GetName() string {
	return agg.name
}
