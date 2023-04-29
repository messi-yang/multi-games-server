package gamerhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) QueryGamers(c *gin.Context) {
	pgUow := pguow.NewDummyUow()

	gamerAppService := provideGamerAppService(pgUow)
	gamerDtos, err := gamerAppService.QueryGamers(gamerappsrv.QueryGamersQuery{})
	if err != nil {
		pgUow.Rollback()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.Commit()
	c.JSON(http.StatusOK, queryGamersReponse(gamerDtos))
}
