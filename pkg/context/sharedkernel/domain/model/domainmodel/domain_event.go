package domainmodel

import "time"

type DomainEvent interface {
	GetName() string
	GetOccurredOn() time.Time
}
