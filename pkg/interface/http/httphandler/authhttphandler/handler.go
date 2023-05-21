package authhttphandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GoToGoogleAuthUrl(c *gin.Context) {
	googleAuthInfraService := provideGoogleAuthInfraService()

	authUrl := googleAuthInfraService.GenerateAuthUrl(googleauthinfrasrv.GenerateAuthUrlCommand{})
	c.Redirect(http.StatusFound, authUrl)
}

func (httpHandler *HttpHandler) HandleGoogleAuthCallback(c *gin.Context) {
	code := c.Query("code")
	googleAuthInfraService := provideGoogleAuthInfraService()
	userEmailAddress, err := googleAuthInfraService.GetUserEmailAddress(googleauthinfrasrv.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		return
	}

	pgUow := pguow.NewUow()

	identityAppService := provideIdentityAppService(pgUow)

	userDto, userFound, err := identityAppService.FindUserByEmailAddress(identityappsrv.FindUserByEmailAddressQuery{
		EmailAddress: userEmailAddress,
	})
	if err != nil {
		pgUow.RevertChanges()
		return
	}

	var userIdDto uuid.UUID
	if userFound {
		userIdDto = userDto.Id
	} else {
		userIdDto, err = identityAppService.Register(identityappsrv.RegisterCommand{EmailAddress: userEmailAddress})
		if err != nil {
			pgUow.RevertChanges()
			return
		}
	}

	accessToken, err := identityAppService.Login(
		identityappsrv.LoginCommand{UserId: userIdDto},
	)
	if err != nil {
		pgUow.RevertChanges()
		return
	}

	pgUow.SaveChanges()

	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/auth/sign-in-success/?access_token=%v", clientUrl, accessToken))
}
