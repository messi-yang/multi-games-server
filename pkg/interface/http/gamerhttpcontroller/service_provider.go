package gamerhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/pgrepository"
)

func provideGamerAppService() (userAppService gamerappservice.Service, err error) {
	gamerRepository, err := pgrepository.NewGamerRepository()
	if err != nil {
		return userAppService, err
	}
	return gamerappservice.NewService(gamerRepository), nil
}
