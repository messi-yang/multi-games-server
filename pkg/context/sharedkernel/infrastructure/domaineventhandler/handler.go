package domaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/domaineventhandler/userdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

func Run() {
	domainEventDispatcher := pguow.GetDomainEventDispatcher()

	domainEventDispatcher.Register(usermodel.UserCreated{}, userdomaineventhandler.NewUserCreatedHandler())
}
