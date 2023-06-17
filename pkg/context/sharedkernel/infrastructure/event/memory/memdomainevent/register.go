package memdomainevent

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type Register struct {
	mediator *Mediator
}

func NewRegister() *Register {
	return &Register{
		mediator: GetMediator(),
	}
}

func (register *Register) Register(domainEvent domain.DomainEvent, handler Handler) {
	register.mediator.Register(domainEvent, handler)
}
