package httprouter

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
)

func provideIdentityAppService(uow pguow.Uow) identityappsrv.Service {
	userRepo := pgrepo.NewUserRepo(uow)
	identityService := identitydomainsrv.NewService(userRepo)
	return identityappsrv.NewService(userRepo, identityService, os.Getenv("AUTH_SECRET"))
}
