package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

type GameApplicationService struct {
	GameService *service.GameService
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
	gameService, _ := service.NewGameService(service.WithPostgresGameRepository())
	return func(service *GameApplicationService) error {
		service.GameService = gameService
		return nil
	}
}

func (service *GameApplicationService) GetFirstGameId() (valueobject.GameId, error) {
	games, err := service.GameService.GeAllGames()
	if err != nil {
		return valueobject.GameId{}, err
	}
	return games[0].GetId(), nil
}
