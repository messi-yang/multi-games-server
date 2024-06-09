package domain

type DomainEventDispatchableAggregate interface {
	Aggregate
	PopDomainEvents() []DomainEvent
}
