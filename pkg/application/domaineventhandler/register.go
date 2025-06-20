package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
)

func RegisterEvents() {
	domainEventRegister := memdomaineventhandler.NewRegister()
	domainEventRegister.Register(roommodel.RoomCreated{}, ProvideRoomCreatedHandler())
	domainEventRegister.Register(roommodel.RoomDeleted{}, ProvideRoomDeletedHandler())
}
