package itemhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappsrv"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	itemAppService itemappsrv.Service
}

var httpHandlerSingleton *HttpHandler

func NewHttpHandler(
	itemAppService itemappsrv.Service,
) *HttpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &HttpHandler{itemAppService: itemAppService}
}

func (httpHandler *HttpHandler) QueryItems(c *gin.Context) {
	itemDtos, err := httpHandler.itemAppService.QueryItems(itemappsrv.QueryItemsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryItemsReponseDto(itemDtos))
}
