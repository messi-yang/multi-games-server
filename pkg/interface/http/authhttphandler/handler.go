package authhttphandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/service/identityappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleauthinfraservice"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type httpHandler struct {
	googleAuthInfraService googleauthinfraservice.Service
	identityAppService     identityappservice.Service
	gamerappservice        gamerappservice.Service
}

var httpHandlerSingleton *httpHandler

func newHttpHandler(
	googleAuthInfraService googleauthinfraservice.Service,
	identityAppService identityappservice.Service,
	gamerappservice gamerappservice.Service,
) *httpHandler {
	if httpHandlerSingleton != nil {
		return httpHandlerSingleton
	}
	return &httpHandler{googleAuthInfraService, identityAppService, gamerappservice}
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

	userDto, userFound, err := httpHandler.identityAppService.FindUserByEmailAddress(identityappservice.FindUserByEmailAddressQuery{
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
		newUserIdDto, err := httpHandler.identityAppService.Register(identityappservice.RegisterCommand{EmailAddress: userEmailAddress})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		userIdDto = newUserIdDto

		if _, err = httpHandler.gamerappservice.CreateGamer(gamerappservice.CreateGamerCommand{
			UserId: userIdDto,
		}); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Registred")
	}

	accessToken, err := httpHandler.identityAppService.Login(
		identityappservice.LoginCommand{UserId: userIdDto},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/?access_token=%v", clientUrl, accessToken))
}
