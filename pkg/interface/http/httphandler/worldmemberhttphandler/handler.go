package worldmemberhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetWorldMembers(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)
	worldIdDto, err := uuid.Parse(c.Param("worldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewDummyUow()
	worldAccessAppService := providedependency.ProvideWorldAccessAppService(pgUow)

	worldMemberDto, err := worldAccessAppService.GetUserWorldMember(worldaccessappsrv.GetUserWorldMemberQuery{
		WorldId: worldIdDto,
		UserId:  userIdDto,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if worldMemberDto == nil {
		c.String(http.StatusForbidden, "you're not permitted to do this")
		return
	}

	worldMemberDtos, err := worldAccessAppService.GetWorldMembers(worldaccessappsrv.GetWorldMembersQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorldMembersResponse(worldMemberDtos))
}
