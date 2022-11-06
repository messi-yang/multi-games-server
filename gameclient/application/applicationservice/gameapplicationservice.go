package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
)

type GameApplicationService struct {
	GameService *gameservice.GameService
}

type gameApplicationServiceConfiguration func(service *GameApplicationService) error

func NewGameApplicationService(cfgs ...gameApplicationServiceConfiguration) (*GameApplicationService, error) {
	service := &GameApplicationService{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithGameService() gameApplicationServiceConfiguration {
	gameService, _ := gameservice.NewGameService()
	return func(service *GameApplicationService) error {
		service.GameService = gameService
		return nil
	}
}

func (service *GameApplicationService) GetAllGames() []game.Game {
	return service.GameService.GetAllGames()
}
