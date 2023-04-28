package domainmodel

import "time"

type DomainEvent interface {
	GetOccurredOn() time.Time
}
