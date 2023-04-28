package domainmodel

type Aggregate interface {
	AddDomainEvent(DomainEvent)
	GetDomainEvents() []DomainEvent
}
