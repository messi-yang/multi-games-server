package worldapi

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
)

func newWorldAppService(presenter worldappservice.Presenter) (worldAppService worldappservice.Service, err error) {
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
	return worldappservice.NewService(worldRepository, unitRepository, itemRepository, presenter), nil
}
