package authhttphandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/service/identityappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleauthinfraservice"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	googleAuthInfraService googleauthinfraservice.Service
	identityAppService     identityappservice.Service
}

var httpHandlerSingleton *httpHandler

func newHttpHandler(
	googleAuthInfraService googleauthinfraservice.Service, identityAppService identityappservice.Service,
) *httpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &httpHandler{googleAuthInfraService: googleAuthInfraService, identityAppService: identityAppService}
}

func (httpHandler *httpHandler) goToGoogleAuthUrl(c *gin.Context) {
	authUrl := httpHandler.googleAuthInfraService.GenerateAuthUrl(googleauthinfraservice.GenerateAuthUrlCommand{})
	c.Redirect(http.StatusFound, authUrl)
}

func (httpHandler *httpHandler) googleAuthCallback(c *gin.Context) {
	code := c.Query("code")
	userEmailAddress, err := httpHandler.googleAuthInfraService.GetUserEmailAddress(googleauthinfraservice.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	accessToken, err := httpHandler.identityAppService.LoginOrRegister(
		identityappservice.LoginOrRegisterCommand{EmailAddress: userEmailAddress},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/?access_token=%v", clientUrl, accessToken))
}
