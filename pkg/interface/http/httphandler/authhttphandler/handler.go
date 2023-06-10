package authhttphandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/authappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/userappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GoogleOauthState struct {
	ClientPath string `json:"clientPath"`
}

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GoToGoogleAuthUrl(c *gin.Context) {
	clientPath := c.Query("client_path")
	stateInBytes, err := json.Marshal(GoogleOauthState{
		ClientPath: clientPath,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	state := string(stateInBytes)

	googleAuthInfraService := providedependency.ProvideGoogleAuthInfraService()
	authUrl, err := googleAuthInfraService.GenerateAuthUrl(googleauthinfrasrv.GenerateAuthUrlCommand{
		State: state,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Redirect(http.StatusFound, authUrl)
}

func (httpHandler *HttpHandler) HandleGoogleAuthCallback(c *gin.Context) {
	code := c.Query("code")
	state, err := jsonutil.Unmarshal[GoogleOauthState]([]byte(c.Query("state")))
	if err != nil {
		return
	}

	googleAuthInfraService := providedependency.ProvideGoogleAuthInfraService()
	userEmailAddress, err := googleAuthInfraService.GetUserEmailAddress(googleauthinfrasrv.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		return
	}

	pgUow := pguow.NewUow()

	userAppService := providedependency.ProvideUserAppService(pgUow)
	authAppService := providedependency.ProvideAuthAppService(pgUow)

	userDto, userFound, err := userAppService.FindUserByEmailAddress(userappsrv.FindUserByEmailAddressQuery{
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
		userIdDto, err = authAppService.Register(authappsrv.RegisterCommand{EmailAddress: userEmailAddress})
		if err != nil {
			pgUow.RevertChanges()
			return
		}
	}

	accessToken, err := authAppService.Login(
		authappsrv.LoginCommand{UserId: userIdDto},
	)
	if err != nil {
		pgUow.RevertChanges()
		return
	}

	pgUow.SaveChanges()

	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(
		http.StatusFound,
		fmt.Sprintf(
			"%s/auth/sign-in-success/?access_token=%v&client_path=%v",
			clientUrl,
			accessToken,
			state.ClientPath,
		),
	)
}
