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

	userDto, err := userAppService.GetUser(userappsrv.GetUserQuery{
		UserId: userIdDto,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getMyUserResponse(userDto))
}

func (httpHandler *HttpHandler) UpdateMyUser(c *gin.Context) {
	userIdDto := httputil.GetUserId(c)

	var requestBody updateMyUserRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	userAppService := providedependency.ProvideUserAppService(pgUow)

	if err := userAppService.UpdateUser(userappsrv.UpdateUserCommand{
		UserId:   userIdDto,
		Username: requestBody.Username,
	}); err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	updatedUserDto, err := userAppService.GetUser(userappsrv.GetUserQuery{
		UserId: userIdDto,
	})
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow.SaveChanges()
	c.JSON(http.StatusOK, updateMyUserResponse(updatedUserDto))
}
