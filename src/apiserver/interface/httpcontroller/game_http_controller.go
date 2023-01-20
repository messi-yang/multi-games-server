package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/appservice"
	"github.com/gin-gonic/gin"
)

type GameHttpController struct {
	itemAppService appservice.GameAppService
}

func NewGameHttpController(
	itemAppService appservice.GameAppService,
) *GameHttpController {
	return &GameHttpController{
		itemAppService: itemAppService,
	}
}

func (controller *GameHttpController) GetAllHandler(c *gin.Context) {
	controller.itemAppService.GetAll(NewHttpPresenter(c))
}
