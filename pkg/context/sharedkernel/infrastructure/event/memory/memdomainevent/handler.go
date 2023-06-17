package memdomainevent

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

type Handler interface {
	Handle(pguow.Uow, domain.DomainEvent) error
}
