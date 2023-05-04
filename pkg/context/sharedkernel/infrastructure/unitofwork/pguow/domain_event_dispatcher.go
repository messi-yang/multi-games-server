package pguow

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"

type DomainEventDispatcher interface {
	Register(domain.DomainEvent, DomainEventHandler)
	Dispatch(Uow, domain.DomainEvent) error
}

type dispatch struct {
	domainEventHandlersMap map[string][]DomainEventHandler
}

var dispatchSingleton *dispatch

func GetDomainEventDispatcher() DomainEventDispatcher {
	if dispatchSingleton != nil {
		return dispatchSingleton
	}
	dispatchSingleton = &dispatch{
		domainEventHandlersMap: make(map[string][]DomainEventHandler),
	}
	return dispatchSingleton
}

func (dispatch *dispatch) Register(domainEvent domain.DomainEvent, domainEventHandler DomainEventHandler) {
	if dispatch.domainEventHandlersMap[domainEvent.GetEventName()] == nil {
		dispatch.domainEventHandlersMap[domainEvent.GetEventName()] = []DomainEventHandler{}
	}
	domainEventHandlers := dispatch.domainEventHandlersMap[domainEvent.GetEventName()]
	domainEventHandlers = append(
		domainEventHandlers,
		domainEventHandler,
	)
	dispatch.domainEventHandlersMap[domainEvent.GetEventName()] = domainEventHandlers
}

func (dispatch *dispatch) Dispatch(uow Uow, domainEvent domain.DomainEvent) error {
	domainEventHandlers := dispatch.domainEventHandlersMap[domainEvent.GetEventName()]
	if domainEventHandlers == nil {
		return nil
	}
	for _, domainEventHandler := range domainEventHandlers {
		err := domainEventHandler.Handle(uow, domainEvent)
		if err != nil {
			return err
		}
	}
	return nil
}
