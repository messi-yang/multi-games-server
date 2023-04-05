package worldhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres"
)

func provideWorldAppService() (worldAppService worldappservice.Service, err error) {
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		return worldAppService, err
	}
	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		return worldAppService, err
	}
	unitRepository, err := postgres.NewUnitRepository()
	if err != nil {
		return worldAppService, err
	}
	return worldappservice.NewService(worldRepository, unitRepository, itemRepository), nil
}
