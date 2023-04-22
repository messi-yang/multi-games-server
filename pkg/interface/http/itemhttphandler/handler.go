package itemhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	itemAppService itemappservice.Service
}

var httpHandlerSingleton *httpHandler

func newHttpHandler(
	itemAppService itemappservice.Service,
) *httpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &httpHandler{itemAppService: itemAppService}
}

func (httpHandler *httpHandler) queryItems(c *gin.Context) {
	itemDtos, err := httpHandler.itemAppService.QueryItems(itemappservice.QueryItemsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryItemsReponseDto(itemDtos))
}
