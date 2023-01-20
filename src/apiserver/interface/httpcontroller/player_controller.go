package httpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/playerappservice"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	playerAppService playerappservice.Service
}

func NewPlayerController(
	playerAppService playerappservice.Service,
) *PlayerController {
	return &PlayerController{
		playerAppService: playerAppService,
	}
}

func (controller *PlayerController) GetAllHandler(c *gin.Context) {
	controller.playerAppService.GetAllPlayers(NewHttpPresenter(c))
}
