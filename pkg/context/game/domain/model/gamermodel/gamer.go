package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Gamer struct {
	id                   commonmodel.GamerId
	userId               sharedkernelmodel.UserId
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Gamer)(nil)

func NewGamer(
	id commonmodel.GamerId,
	userId sharedkernelmodel.UserId,
) Gamer {
	return Gamer{
		id:                   id,
		userId:               userId,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (gamer *Gamer) PopDomainEvents() []domain.DomainEvent {
	return gamer.domainEventCollector.PopAll()
}

func (gamer *Gamer) GetId() commonmodel.GamerId {
	return gamer.id
}

func (gamer *Gamer) GetUserId() sharedkernelmodel.UserId {
	return gamer.userId
}
