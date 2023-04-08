package userhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/userappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/pgrepository"
)

func provideUserAppService() (userAppService userappservice.Service, err error) {
	userRepository, err := pgrepository.NewUserRepository()
	if err != nil {
		return userAppService, err
	}
	return userappservice.NewService(userRepository), nil
}
