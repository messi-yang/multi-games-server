package domain

type DomainEventCollector struct {
	domainEvents []DomainEvent
}

func NewDomainEventCollector() *DomainEventCollector {
	return &DomainEventCollector{
		domainEvents: make([]DomainEvent, 0),
	}
}

func (collector *DomainEventCollector) Add(domainEvent DomainEvent) {
	collector.domainEvents = append(collector.domainEvents, domainEvent)
}

func (collector *DomainEventCollector) PopAll() []DomainEvent {
	poppedDomainEvents := collector.domainEvents
	collector.domainEvents = make([]DomainEvent, 0)
	return poppedDomainEvents
}
