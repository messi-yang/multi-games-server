package domain

type DomainEventDispatcher interface {
	Dispatch(Aggregate) error
}
