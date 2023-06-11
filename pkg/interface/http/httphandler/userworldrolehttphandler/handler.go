package userworldrolehttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetUserWorldRoles(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	worldAccessAppService := providedependency.ProvideWorldAccessAppService(pgUow)

	_, userWorldRoleFound, err := worldAccessAppService.FindUserWorldRole(worldaccessappsrv.FindUserWorldRoleQuery{
		WorldId: worldIdDto,
		UserId:  userIdDto,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if !userWorldRoleFound {
		c.String(http.StatusForbidden, "you're not permitted to do this")
		return
	}

	userWorldRoleDtos, err := worldAccessAppService.GetUserWorldRoles(worldaccessappsrv.GetUserWorldRolesQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUserWorldRolesResponse(userWorldRoleDtos))
}
