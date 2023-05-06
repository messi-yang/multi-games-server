package userdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

type UserCreatedHandler struct {
}

func NewUserCreatedHandler() pguow.DomainEventHandler {
	return &UserCreatedHandler{}
}

func (handler UserCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	userCreated := domainEvent.(usermodel.UserCreated)
	gamerAppService := provideGamerAppService(uow)

	if _, err := gamerAppService.CreateGamer(gamerappsrv.CreateGamerCommand{
		UserId: userCreated.GetUserId().Uuid(),
	}); err != nil {
		return err
	}
	return nil
}
