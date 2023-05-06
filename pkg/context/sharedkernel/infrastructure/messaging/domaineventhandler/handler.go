package domaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/domaineventhandler/userdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memdomainevent"
)

func Run() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(usermodel.UserCreated{}, userdomaineventhandler.NewUserCreatedHandler())
}
