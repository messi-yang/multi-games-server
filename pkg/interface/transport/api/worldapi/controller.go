package worldapi

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/api"
	"github.com/gin-gonic/gin"
)

func QueryWorldHandler(c *gin.Context) {
	httpPresenter := api.NewHttpPresenter(c)
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		httpPresenter.OnError(err)
		return
	}
	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		httpPresenter.OnError(err)
		return
	}
	unitRepository, err := postgres.NewUnitRepository()
	if err != nil {
		httpPresenter.OnError(err)
		return
	}
	worldAppService := worldappservice.NewService(worldRepository, unitRepository, itemRepository, httpPresenter)

	worldAppService.QueryWorlds()
}

func CreateWorldHandler(c *gin.Context) {
	var requestDto worldappservice.CreateWorldRequestDto

	if err := c.BindJSON(&requestDto); err != nil {
		return
	}
	fmt.Println(requestDto)

	httpPresenter := api.NewHttpPresenter(c)
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		httpPresenter.OnError(err)
		return
	}
	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		httpPresenter.OnError(err)
		return
	}
	unitRepository, err := postgres.NewUnitRepository()
	if err != nil {
		httpPresenter.OnError(err)
		return
	}
	worldAppService := worldappservice.NewService(worldRepository, unitRepository, itemRepository, httpPresenter)

	worldAppService.CreateWorld(requestDto.UserId)
}
