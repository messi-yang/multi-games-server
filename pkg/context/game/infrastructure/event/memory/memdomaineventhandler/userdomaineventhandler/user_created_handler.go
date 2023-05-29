package userdomaineventhandler

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

type UserCreatedHandler struct {
}

func NewUserCreatedHandler() memdomainevent.Handler {
	return &UserCreatedHandler{}
}

func (handler UserCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	fmt.Println(domainEvent)
	userCreated := domainEvent.(sharedkernelmodel.UserCreated)
	gamerAppService := provideGamerAppService(uow)

	if _, err := gamerAppService.CreateGamer(gamerappsrv.CreateGamerCommand{
		UserId: userCreated.GetUserId().Uuid(),
	}); err != nil {
		return err
	}
	return nil
}
