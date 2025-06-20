package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
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
	roomCreated := domainEvent.(roommodel.RoomCreated)

	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameAccountRepo := pgrepo.NewGameAccountRepo(uow, domainEventDispatcher)

	gameAccount, err := gameAccountRepo.GetGameAccountOfUser(roomCreated.GetUserId())
	if err != nil {
		return err
	}
	gameAccount.AddRoomsCount()
	err = gameAccountRepo.Update(gameAccount)
	if err != nil {
		return err
	}

	gameRepo := pgrepo.NewGameRepo(uow, domainEventDispatcher)
	return gameRepo.Add(gamemodel.NewGame(
		roomCreated.GetRoomId(),
		"hello_world",
		map[string]interface{}{},
	))
}
