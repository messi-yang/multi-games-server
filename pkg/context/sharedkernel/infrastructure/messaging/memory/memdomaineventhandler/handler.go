package memdomaineventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memory/memdomaineventhandler/userdomaineventhandler"
)

func Run() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(usermodel.UserCreated{}, userdomaineventhandler.NewUserCreatedHandler())
}
