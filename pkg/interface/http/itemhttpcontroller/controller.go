package itemhttpcontroller

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httpdto"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func queryWorldHandler(c *gin.Context) {
	itemAppService, err := provideItemAppService()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	items, err := itemAppService.GetItems(itemappservice.GetItemsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDtos := lo.Map(items, func(item itemmodel.ItemAgg, _ int) httpdto.ItemAggDto {
		return httpdto.NewItemAggDto(item)
	})
	c.JSON(http.StatusOK, getItemsReponseDto(itemDtos))
}
