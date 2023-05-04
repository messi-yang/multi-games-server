package domain

type Aggregate interface {
	PopDomainEvents() []DomainEvent
}
