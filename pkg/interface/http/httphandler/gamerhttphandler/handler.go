package gamerhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) QueryGamers(c *gin.Context) {
	pgUow := pguow.NewDummyUow()

	gamerAppService := providedependency.ProvideGamerAppService(pgUow)
	gamerDtos, err := gamerAppService.QueryGamers(gamerappsrv.QueryGamersQuery{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, queryGamersReponse(gamerDtos))
}
