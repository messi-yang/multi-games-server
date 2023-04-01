package itemhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
)

func provideItemAppService() (itemAppService itemappservice.Service, err error) {
	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		return itemAppService, err
	}
	return itemappservice.NewService(itemRepository), nil
}
