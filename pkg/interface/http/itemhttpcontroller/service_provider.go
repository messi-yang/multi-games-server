package itemhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/pgrepository"
)

func provideItemAppService() (itemAppService itemappservice.Service, err error) {
	itemRepository, err := pgrepository.NewItemRepository()
	if err != nil {
		return itemAppService, err
	}
	return itemappservice.NewService(itemRepository), nil
}
