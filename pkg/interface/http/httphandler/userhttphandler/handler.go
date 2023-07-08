package userhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httputil"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetMyUser(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	pgUow := pguow.NewDummyUow()
	userAppService := providedependency.ProvideUserAppService(pgUow)

	userDto, err := userAppService.GetUserQuery(userappsrv.GetUserQuery{
		UserId: userIdDto,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, getMyUserResponse(userDto))
}
