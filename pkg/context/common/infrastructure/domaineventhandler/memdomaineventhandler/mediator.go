package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
)

type Mediator struct {
	handlersMap map[string][]Handler
}

var mediatorSingleton *Mediator

func GetMediator() *Mediator {
	if mediatorSingleton != nil {
		return mediatorSingleton
	}
	mediatorSingleton = &Mediator{
		handlersMap: make(map[string][]Handler),
	}
	return mediatorSingleton
}

func (mediator *Mediator) Register(domainEvent domain.DomainEvent, newHandler Handler) {
	if mediator.handlersMap[domainEvent.GetEventName()] == nil {
		mediator.handlersMap[domainEvent.GetEventName()] = make([]Handler, 0)
	}
	handlers := mediator.handlersMap[domainEvent.GetEventName()]
	handlers = append(
		handlers,
		newHandler,
	)
	mediator.handlersMap[domainEvent.GetEventName()] = handlers
}

func (mediator *Mediator) Dispatch(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	handlers := mediator.handlersMap[domainEvent.GetEventName()]
	if handlers == nil {
		return nil
	}
	for _, handler := range handlers {
		err := handler.Handle(uow, domainEvent)
		if err != nil {
			return err
		}
	}
	return nil
}
