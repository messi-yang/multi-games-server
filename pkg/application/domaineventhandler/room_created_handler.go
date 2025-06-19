package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/domainevent"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
)

type RoomCreatedHandler struct{}

func NewRoomCreatedHandler() memdomaineventhandler.Handler {
	return &RoomCreatedHandler{}
}

func ProvideRoomCreatedHandler() memdomaineventhandler.Handler {
	return NewRoomCreatedHandler()
}

func (handler RoomCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	roomCreated := domainEvent.(domainevent.RoomCreated)

	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameAccountRepo := pgrepo.NewGameAccountRepo(uow, domainEventDispatcher)

	gameAccount, err := gameAccountRepo.GetGameAccountOfUser(roomCreated.GetUserId())
	if err != nil {
		return err
	}
	gameAccount.AddRoomsCount()
	return gameAccountRepo.Update(gameAccount)
}
