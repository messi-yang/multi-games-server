package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/playerappservice"
	"github.com/gin-gonic/gin"
)

type PlayerHttpController struct {
	playerAppService playerappservice.Service
}

func NewPlayerHttpController(
	playerAppService playerappservice.Service,
) *PlayerHttpController {
	return &PlayerHttpController{
		playerAppService: playerAppService,
	}
}

func (controller *PlayerHttpController) GetAllHandler(c *gin.Context) {
	controller.playerAppService.GetAllPlayers(NewHttpPresenter(c))
}
