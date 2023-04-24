package authhttphandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct {
	googleAuthInfraService googleauthinfrasrv.Service
	identityAppService     identityappsrv.Service
	gamerappsrv            gamerappsrv.Service
}

var httpHandlerSingleton *HttpHandler

func NewHttpHandler(
	googleAuthInfraService googleauthinfrasrv.Service,
	identityAppService identityappsrv.Service,
	gamerappsrv gamerappsrv.Service,
) *HttpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &HttpHandler{googleAuthInfraService, identityAppService, gamerappsrv}
}

func (httpHandler *HttpHandler) GoToGoogleAuthUrl(c *gin.Context) {
	authUrl := httpHandler.googleAuthInfraService.GenerateAuthUrl(googleauthinfrasrv.GenerateAuthUrlCommand{})
	c.Redirect(http.StatusFound, authUrl)
}

func (httpHandler *HttpHandler) HandleGoogleAuthCallback(c *gin.Context) {
	code := c.Query("code")
	userEmailAddress, err := httpHandler.googleAuthInfraService.GetUserEmailAddress(googleauthinfrasrv.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userDto, userFound, err := httpHandler.identityAppService.FindUserByEmailAddress(identityappsrv.FindUserByEmailAddressQuery{
		EmailAddress: userEmailAddress,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var userIdDto uuid.UUID
	if userFound {
		fmt.Println("Logged in")
		userIdDto = userDto.Id
	} else {
		newUserIdDto, err := httpHandler.identityAppService.Register(identityappsrv.RegisterCommand{EmailAddress: userEmailAddress})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		userIdDto = newUserIdDto

		if _, err = httpHandler.gamerappsrv.CreateGamer(gamerappsrv.CreateGamerCommand{
			UserId: userIdDto,
		}); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Registred")
	}

	accessToken, err := httpHandler.identityAppService.Login(
		identityappsrv.LoginCommand{UserId: userIdDto},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/?access_token=%v", clientUrl, accessToken))
}
