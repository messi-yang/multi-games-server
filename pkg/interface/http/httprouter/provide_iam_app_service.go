package httprouter

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/identitymodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/event/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func provideIdentityAppService(uow pguow.Uow) identityappsrv.Service {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)
	identityService := identitymodel.NewIdentityService(userRepo)
	return identityappsrv.NewService(userRepo, identityService, os.Getenv("AUTH_SECRET"))
}
