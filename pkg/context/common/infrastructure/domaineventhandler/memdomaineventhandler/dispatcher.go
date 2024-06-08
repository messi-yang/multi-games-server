package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
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

func (dispatch *Dispatch) Dispatch(aggregate domain.DomainEventDispatchableAggregate) error {
	domainEvents := aggregate.PopDomainEvents()
	for _, domainEvent := range domainEvents {
		err := dispatch.mediator.Dispatch(dispatch.uow, domainEvent)
		if err != nil {
			return err
		}
	}
	return nil
}
