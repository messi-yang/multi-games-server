package authhttphandler

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/postgres/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memory/memdomainevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func provideGoogleAuthInfraService() googleauthinfrasrv.Service {
	return googleauthinfrasrv.NewService()
}

func provideIdentityAppService(uow pguow.Uow) identityappsrv.Service {
	userRepo := pgrepo.NewUserRepo(uow)
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	identityDomainService := identitydomainsrv.NewService(userRepo, domainEventDispatcher)
	return identityappsrv.NewService(userRepo, identityDomainService, os.Getenv("AUTH_SECRET"))
}
