package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type GamerAgg struct {
	id           commonmodel.GamerIdVo
	userId       sharedkernelmodel.UserIdVo
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*GamerAgg)(nil)

func NewGamerAgg(
	id commonmodel.GamerIdVo,
	userId sharedkernelmodel.UserIdVo,
) GamerAgg {
	return GamerAgg{id: id, userId: userId, domainEvents: []domainmodel.DomainEvent{}}
}

func (agg *GamerAgg) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	agg.domainEvents = append(agg.domainEvents, domainEvent)
}

func (agg *GamerAgg) GetDomainEvents() []domainmodel.DomainEvent {
	return agg.domainEvents
}

func (agg *GamerAgg) GetId() commonmodel.GamerIdVo {
	return agg.id
}

func (agg *GamerAgg) GetUserId() sharedkernelmodel.UserIdVo {
	return agg.userId
}
