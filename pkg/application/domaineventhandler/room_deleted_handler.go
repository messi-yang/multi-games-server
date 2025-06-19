package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/domainevent"
)

type RoomDeletedHandler struct{}

func NewRoomDeletedHandler() memdomaineventhandler.Handler {
	return &RoomDeletedHandler{}
}

func ProvideRoomDeletedHandler() memdomaineventhandler.Handler {
	return NewRoomDeletedHandler()
}

func (handler RoomDeletedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	roomDeleted := domainEvent.(domainevent.RoomDeleted)

	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameAccountRepo := pgrepo.NewGameAccountRepo(uow, domainEventDispatcher)

	gameAccount, err := gameAccountRepo.GetGameAccountOfUser(roomDeleted.GetUserId())
	if err != nil {
		return err
	}
	gameAccount.SubtractRoomsCount()
	return gameAccountRepo.Update(gameAccount)
}
