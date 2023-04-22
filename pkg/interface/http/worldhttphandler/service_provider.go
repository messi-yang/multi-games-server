package worldhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepository"
)

func provideWorldAppService() (worldAppService worldappservice.Service, err error) {
	worldRepository, err := pgrepository.NewWorldRepository()
	if err != nil {
		return worldAppService, err
	}
	itemRepository, err := pgrepository.NewItemRepository()
	if err != nil {
		return worldAppService, err
	}
	unitRepository, err := pgrepository.NewUnitRepository()
	if err != nil {
		return worldAppService, err
	}
	return worldappservice.NewService(worldRepository, unitRepository, itemRepository), nil
}
