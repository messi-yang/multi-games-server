package userhttphandler

import (
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GetMyUser(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	pgUow := pguow.NewDummyUow()
	getUserUseCase := usecase.ProvideGetUserUseCase(pgUow)
	userDto, err := getUserUseCase.Execute(*authorizedUserIdDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func (httpHandler *HttpHandler) UpdateMyUser(c *gin.Context) {
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	if authorizedUserIdDto == nil {
		c.String(http.StatusUnauthorized, "not authorized")
		return
	}

	var requestBody updateMyUserRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pgUow := pguow.NewUow()
	updateUserUseCase := usecase.ProvideUpdateUserUseCase(pgUow)
	updatedUserDto, err := updateUserUseCase.Execute(*authorizedUserIdDto, requestBody.Username, requestBody.FriendlyName)
	if err != nil {
		pgUow.RevertChanges()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pgUow.SaveChanges()

	c.JSON(http.StatusOK, updatedUserDto)
}
