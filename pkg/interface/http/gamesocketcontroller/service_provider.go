package gamesocketcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres"
)

func provideGameAppService() (gameAppService gameappservice.Service, err error) {
	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		return gameAppService, err
	}
	playerRepository := memrepo.NewPlayerMemRepository()
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		return gameAppService, err
	}
	unitRepository, err := postgres.NewUnitRepository()
	if err != nil {
		return gameAppService, err
	}
	gameService := service.NewGameService(worldRepository, playerRepository, unitRepository, itemRepository)
	return gameappservice.NewService(
		worldRepository, playerRepository, unitRepository, itemRepository, gameService,
	), nil
}
