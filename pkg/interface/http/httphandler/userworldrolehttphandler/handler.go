package userworldrolehttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/accessappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetUserWorldRoles(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	accessAppService := providedependency.ProvideAccessAppService(pgUow)

	userWorldRoleDtos, err := accessAppService.GetUserWorldRoles(accessappsrv.GetUserWorldRolesQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserWorldRolesResponse(userWorldRoleDtos))
}
