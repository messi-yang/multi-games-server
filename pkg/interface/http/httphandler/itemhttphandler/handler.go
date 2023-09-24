package itemhttphandler

import (
	"net/http"
	"strings"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	itemIdsQueryString := c.Request.URL.Query().Get("itemIds")
	itemIdStrings := strings.Split(itemIdsQueryString, ",")
	itemIdDtos, err := commonutil.MapWithError(itemIdStrings, func(_ int, itemIdString string) (uuid.UUID, error) {
		return uuid.Parse(itemIdString)
	})

	pgUow := pguow.NewDummyUow()

	itemAppService := providedependency.ProvideItemAppService(pgUow)
	itemDtos, err := itemAppService.GetItemsOfIds(itemappsrv.GetItemsOfIdsQuery{
		ItemIds: itemIdDtos,
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
