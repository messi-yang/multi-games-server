package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/appservice"
	"github.com/gin-gonic/gin"
)

type ItemHttpController struct {
	itemAppService appservice.ItemAppService
}

func NewItemHttpController(
	itemAppService appservice.ItemAppService,
) *ItemHttpController {
	return &ItemHttpController{
		itemAppService: itemAppService,
	}
}

func (controller *ItemHttpController) GetAllHandler(c *gin.Context) {
	controller.itemAppService.GetAllItems(NewHttpPresenter(c))
}
