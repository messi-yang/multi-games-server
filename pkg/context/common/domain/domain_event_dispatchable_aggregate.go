package domain

type DomainEventDispatchableAggregate interface {
	PopDomainEvents() []DomainEvent
}
