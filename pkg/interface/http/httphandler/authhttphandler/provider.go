package authhttphandler

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	iam_pg_repo "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
)

func provideGoogleAuthInfraService() googleauthinfrasrv.Service {
	return googleauthinfrasrv.NewService()
}

func provideGamerAppService(pgUow *pguow.Uow) gamerappsrv.Service {
	gamerRepo := pgrepo.NewGamerRepo(pgUow)
	return gamerappsrv.NewService(gamerRepo)
}

func provideIdentityAppService(pgUow *pguow.Uow) identityappsrv.Service {
	userRepo := iam_pg_repo.NewUserRepo(pgUow)
	identityDomainService := identitydomainsrv.NewService(userRepo)
	return identityappsrv.NewService(userRepo, identityDomainService, os.Getenv("AUTH_SECRET"))
}
