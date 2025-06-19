package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/domainevent"
)

func RegisterEvents() {
	domainEventRegister := memdomaineventhandler.NewRegister()
	domainEventRegister.Register(domainevent.RoomCreated{}, ProvideRoomCreatedHandler())
	domainEventRegister.Register(domainevent.RoomDeleted{}, ProvideRoomDeletedHandler())
}
