package worldaccounthttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldaccountappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) QueryWorldAccounts(c *gin.Context) {
	pgUow := pguow.NewDummyUow()

	worldAccountAppService := providedependency.ProvideWorldAccountAppService(pgUow)
	worldAccountDtos, err := worldAccountAppService.QueryWorldAccounts(worldaccountappsrv.QueryWorldAccountsQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryWorldAccountsReponse(worldAccountDtos))
}
