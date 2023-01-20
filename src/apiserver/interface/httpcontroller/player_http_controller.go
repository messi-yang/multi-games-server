package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/appservice"
	"github.com/gin-gonic/gin"
)

type PlayerHttpController struct {
	playerAppService appservice.PlayerAppService
}

func NewPlayerHttpController(
	playerAppService appservice.PlayerAppService,
) *PlayerHttpController {
	return &PlayerHttpController{
		playerAppService: playerAppService,
	}
}

func (controller *PlayerHttpController) GetAllHandler(c *gin.Context) {
	controller.playerAppService.GetAllPlayers(NewHttpPresenter(c))
}
