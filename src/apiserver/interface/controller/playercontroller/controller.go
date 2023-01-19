package playercontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/playerappservice"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	playerAppService playerappservice.Service
}

func New(
	playerAppService playerappservice.Service,
) *Controller {
	return &Controller{
		playerAppService: playerAppService,
	}
}

func (controller *Controller) GetAllHandler(c *gin.Context) {
	controller.playerAppService.GetAllPlayers(newGinPresenter(c))
}
