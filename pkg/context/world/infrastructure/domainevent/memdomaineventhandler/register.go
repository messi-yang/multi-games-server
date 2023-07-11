package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

func RegisterEvents() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(sharedkernelmodel.UserCreated{}, NewUserCreatedHandler())
	domainEventRegister.Register(sharedkernelmodel.WorldCreated{}, NewWorldCreatedHandler())
}
