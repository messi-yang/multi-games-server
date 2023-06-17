package itemhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, queryItemsReponse(itemDtos))
}
