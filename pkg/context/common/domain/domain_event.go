package domain

import "time"

type DomainEvent interface {
	GetEventName() string
	GetOccurredOn() time.Time
}
