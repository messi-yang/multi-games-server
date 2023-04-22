package gamerhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappservice"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	gamerAppService gamerappservice.Service
}

var httpHandlerSingleton *httpHandler

func newHttpHandler(
	gamerAppService gamerappservice.Service,
) *httpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &httpHandler{gamerAppService: gamerAppService}
}

func (httpHandler *httpHandler) queryGamers(c *gin.Context) {
	gamerDtos, err := httpHandler.gamerAppService.QueryGamers(gamerappservice.QueryGamersQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryGamersReponseDto(gamerDtos))
}
