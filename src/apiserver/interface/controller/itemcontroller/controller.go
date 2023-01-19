package itemcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	itemAppService itemappservice.Service
}

func New(
	itemAppService itemappservice.Service,
) *Controller {
	return &Controller{
		itemAppService: itemAppService,
	}
}

func (controller *Controller) GetAllHandler(c *gin.Context) {
	controller.itemAppService.GetAllItems(newGinPresenter(c))
}
