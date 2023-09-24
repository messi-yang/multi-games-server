package itemhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) QueryItems(c *gin.Context) {
	pgUow := pguow.NewDummyUow()

	itemAppService := providedependency.ProvideItemAppService(pgUow)
	itemDtos, err := itemAppService.QueryItems(itemappsrv.QueryItemsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemViewModels := lo.Map(itemDtos, func(itemDto dto.ItemDto, _ int) viewmodel.ItemViewModel {
		return viewmodel.ItemViewModel(itemDto)
	})

	c.JSON(http.StatusOK, queryItemsReponse(itemViewModels))
}

func (httpHandler *HttpHandler) GetItemsOfIds(c *gin.Context) {
	var requestBody getItemsOfIdsRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()

	itemAppService := providedependency.ProvideItemAppService(pgUow)
	itemDtos, err := itemAppService.GetItemsOfIds(itemappsrv.GetItemsOfIdsQuery{
		ItemIds: requestBody.ItemIds,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemViewModels := lo.Map(itemDtos, func(itemDto dto.ItemDto, _ int) viewmodel.ItemViewModel {
		return viewmodel.ItemViewModel(itemDto)
	})

	c.JSON(http.StatusOK, queryItemsReponse(itemViewModels))
}
