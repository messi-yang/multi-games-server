package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

type ItemHttpController struct {
	itemAppService itemappservice.Service
}

func NewItemHttpController(
	itemAppService itemappservice.Service,
) *ItemHttpController {
	return &ItemHttpController{
		itemAppService: itemAppService,
	}
}

func (controller *ItemHttpController) GetAllHandler(c *gin.Context) {
	controller.itemAppService.GetAllItems(NewHttpPresenter(c))
}
