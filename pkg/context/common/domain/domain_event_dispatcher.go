package domain

type DomainEventDispatcher interface {
	Dispatch(DomainEventDispatchableAggregate) error
}
