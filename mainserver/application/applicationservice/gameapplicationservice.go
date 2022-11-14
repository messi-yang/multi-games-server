package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/gameservice"
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
	gameService, _ := gameservice.NewGameService(gameservice.WithPostgresGameRepository())
	return func(service *GameApplicationService) error {
		service.GameService = gameService
		return nil
	}
}

func (service *GameApplicationService) GetFirstGameId() (gamemodel.GameId, error) {
	games, err := service.GameService.GeAllGames()
	if err != nil {
		return gamemodel.GameId{}, err
	}
	return games[0].GetId(), nil
}
