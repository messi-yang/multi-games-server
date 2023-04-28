package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Gamer struct {
	id           commonmodel.GamerId
	userId       sharedkernelmodel.UserId
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*Gamer)(nil)

func NewGamer(
	id commonmodel.GamerId,
	userId sharedkernelmodel.UserId,
) Gamer {
	return Gamer{id: id, userId: userId, domainEvents: []domainmodel.DomainEvent{}}
}

func (gamerappsrv *Gamer) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	gamerappsrv.domainEvents = append(gamerappsrv.domainEvents, domainEvent)
}

func (gamerappsrv *Gamer) GetDomainEvents() []domainmodel.DomainEvent {
	return gamerappsrv.domainEvents
}

func (gamerappsrv *Gamer) GetId() commonmodel.GamerId {
	return gamerappsrv.id
}

func (gamerappsrv *Gamer) GetUserId() sharedkernelmodel.UserId {
	return gamerappsrv.userId
}
