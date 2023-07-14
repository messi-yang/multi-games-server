package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

func RegisterEvents() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(worldmodel.WorldCreated{}, NewWorldCreatedHandler())
	domainEventRegister.Register(worldmodel.WorldDeleted{}, NewWorldDeletedHandler())
}
