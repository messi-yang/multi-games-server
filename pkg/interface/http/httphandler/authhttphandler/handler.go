package authhttphandler

import (
	"fmt"
	"net/http"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/gin-gonic/gin"
)

type GoogleOauthState struct {
	ClientRedirectPath string `json:"clientRedirectPath"`
}

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) RedirectToGoogleOauthUrl(c *gin.Context) {
	clientRedirectPath := c.Query("client_redirect_path")

	generateGoogleOauthUrlUseCase := usecase.ProvideGenerateGoogleOauthUrlUseCase()
	googleOauthUrl := generateGoogleOauthUrlUseCase.Execute(clientRedirectPath)

	c.Redirect(http.StatusFound, googleOauthUrl)
}

func (httpHandler *HttpHandler) HandleGoogleOauthCallback(c *gin.Context) {
	fmt.Println("HandleGoogleOauthCallback")
	pgUow := pguow.NewUow()

	loginOrRegisterUserFromGoogleOauthUseCase := usecase.ProvideLoginOrRegisterUserFromGoogleOauthUseCase(pgUow)
	redirectPath, err := loginOrRegisterUserFromGoogleOauthUseCase.Execute(c.Query("code"), c.Query("state"))
	if err != nil {
		pgUow.RevertChanges()
		return
	}
	pgUow.SaveChanges()

	c.Redirect(http.StatusFound, redirectPath)
}
