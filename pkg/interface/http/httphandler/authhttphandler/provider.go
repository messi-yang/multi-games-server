package authhttphandler

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

func provideGoogleAuthInfraService() googleauthinfrasrv.Service {
	return googleauthinfrasrv.NewService()
}

func provideIdentityAppService(uow pguow.Uow) identityappsrv.Service {
	userRepo := pgrepo.NewUserRepo(uow)
	identityDomainService := identitydomainsrv.NewService(userRepo)
	return identityappsrv.NewService(userRepo, identityDomainService, os.Getenv("AUTH_SECRET"))
}
