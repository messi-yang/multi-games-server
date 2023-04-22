package authhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/service/identityappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/persistence/pgrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleauthinfraservice"
)

func provideIdentityAppService() (identityAppService identityappservice.Service, err error) {
	userRepository, err := pgrepository.NewUserRepository()
	if err != nil {
		return identityAppService, err
	}
	identityService := service.NewIdentityService(userRepository)
	return identityappservice.NewService(userRepository, identityService), nil
}

func provideGoogleOauthInfraService() (googleAuthInfraService googleauthinfraservice.Service) {
	return googleauthinfraservice.NewService()
}
