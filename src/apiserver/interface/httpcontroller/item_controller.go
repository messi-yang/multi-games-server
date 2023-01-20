package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemAppService itemappservice.Service
}

func NewItemController(
	itemAppService itemappservice.Service,
) *ItemController {
	return &ItemController{
		itemAppService: itemAppService,
	}
}

func (controller *ItemController) GetAllHandler(c *gin.Context) {
	controller.itemAppService.GetAllItems(NewHttpPresenter(c))
}
