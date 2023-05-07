package memdomainevent

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

type Dispatch struct {
	uow      pguow.Uow
	mediator *Mediator
}

func NewDispatcher(uow pguow.Uow) domain.DomainEventDispatcher {
	return &Dispatch{
		uow:      uow,
		mediator: GetMediator(),
	}
}

func (dispatch *Dispatch) Dispatch(aggregate domain.Aggregate) error {
	domainEvents := aggregate.PopDomainEvents()
	for _, domainEvent := range domainEvents {
		err := dispatch.mediator.Dispatch(dispatch.uow, domainEvent)
		if err != nil {
			return err
		}
	}
	return nil
}
