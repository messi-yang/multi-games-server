package itemcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/common/dto/jsondto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	itemAppService itemappservice.Service
}

func New(
	itemAppService itemappservice.Service,
) *Controller {
	return &Controller{
		itemAppService: itemAppService,
	}
}

func (controller *Controller) HandleGetAllItems(c *gin.Context) {
	items := controller.itemAppService.GetAllItems()
	itemDtos := jsondto.NewItemJsonDtos(items)

	c.JSON(http.StatusOK, itemDtos)
}
