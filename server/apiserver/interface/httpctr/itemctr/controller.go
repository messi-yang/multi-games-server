package itemctr

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemAppService appservice.ItemAppService
}

func NewItemController(
	itemAppService appservice.ItemAppService,
) *ItemController {
	return &ItemController{
		itemAppService: itemAppService,
	}
}

func (controller *ItemController) HandleGetAllItems(c *gin.Context) {
	items := controller.itemAppService.GetAllItems()
	itemDtos := jsondto.NewItemJsonDtos(items)

	c.JSON(http.StatusOK, itemDtos)
}
