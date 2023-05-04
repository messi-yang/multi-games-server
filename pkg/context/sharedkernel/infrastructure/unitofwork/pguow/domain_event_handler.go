package pguow

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/domain"
)

type DomainEventHandler interface {
	Handle(Uow, domain.DomainEvent) error
}
