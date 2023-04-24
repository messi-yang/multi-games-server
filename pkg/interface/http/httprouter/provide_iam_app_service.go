package httprouter

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
)

func provideGoogleAuthInfraService() googleauthinfrasrv.Service {
	return googleauthinfrasrv.NewService()
}

func provideIdentityAppService() (identityappsrv.Service, error) {
	userRepo, err := pgrepo.NewUserRepo()
	if err != nil {
		return nil, err
	}
	identityService := identitydomainsrv.NewService(userRepo)
	return identityappsrv.NewService(userRepo, identityService, os.Getenv("AUTH_SECRET")), nil
}
