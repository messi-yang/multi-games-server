package gamerhttpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappservice"
	"github.com/gin-gonic/gin"
)

func getGamersHandler(c *gin.Context) {
	gamerAppService, err := provideGamerAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	gamerDtos, err := gamerAppService.GetGamers(gamerappservice.GetGamersQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGamersReponseDto(gamerDtos))
}
