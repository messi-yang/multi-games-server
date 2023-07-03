package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
)

func RegisterEvents() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(sharedkernelmodel.UserCreated{}, NewUserCreatedHandler())
}
