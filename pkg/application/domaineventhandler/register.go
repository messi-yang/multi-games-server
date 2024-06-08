package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

func RegisterEvents() {
	domainEventRegister := memdomaineventhandler.NewRegister()
	domainEventRegister.Register(worldmodel.WorldCreated{}, ProvideWorldCreatedHandler())
	domainEventRegister.Register(worldmodel.WorldDeleted{}, ProvideWorldDeletedHandler())
}
