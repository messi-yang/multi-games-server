package gamerhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	gamerAppService gamerappsrv.Service
}

var httpHandlerSingleton *HttpHandler

func NewHttpHandler(
	gamerAppService gamerappsrv.Service,
) *HttpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &HttpHandler{gamerAppService: gamerAppService}
}

func (httpHandler *HttpHandler) QueryGamers(c *gin.Context) {
	gamerDtos, err := httpHandler.gamerAppService.QueryGamers(gamerappsrv.QueryGamersQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryGamersReponseDto(gamerDtos))
}
