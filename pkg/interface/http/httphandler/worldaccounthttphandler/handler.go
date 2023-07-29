package worldaccounthttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldaccountappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
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

	worldAccountViewModels := lo.Map(worldAccountDtos, func(worldAccountDto dto.WorldAccountDto, _ int) viewmodel.WorldAccountViewModel {
		return viewmodel.WorldAccountViewModel(worldAccountDto)
	})

	c.JSON(http.StatusOK, queryWorldAccountsReponse(worldAccountViewModels))
}
