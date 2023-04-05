package itemhttpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/itemappservice"
	"github.com/gin-gonic/gin"
)

func getItemsHandler(c *gin.Context) {
	itemAppService, err := provideItemAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDtos, err := itemAppService.GetItems(itemappservice.GetItemsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getItemsReponseDto(itemDtos))
}
