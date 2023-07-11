package authhttphandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/authappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	iam_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldaccountappsrv"
	world_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GoogleOauthState struct {
	ClientRedirectPath string `json:"clientRedirectPath"`
}

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GoToGoogleAuthUrl(c *gin.Context) {
	clientRedirectPath := c.Query("client_redirect_path")
	stateInBytes := jsonutil.Marshal(GoogleOauthState{
		ClientRedirectPath: clientRedirectPath,
	})
	state := string(stateInBytes)

	googleAuthInfraService := iam_provide_dependency.ProvideGoogleAuthInfraService()
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

	googleAuthInfraService := iam_provide_dependency.ProvideGoogleAuthInfraService()
	userEmailAddress, err := googleAuthInfraService.GetUserEmailAddress(googleauthinfrasrv.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		return
	}

	pgUow := pguow.NewUow()

	userAppService := iam_provide_dependency.ProvideUserAppService(pgUow)
	authAppService := iam_provide_dependency.ProvideAuthAppService(pgUow)
	worldAccountAppService := world_provide_dependency.ProvideWorldAccountAppService(pgUow)

	userDto, err := userAppService.GetUserByEmailAddress(userappsrv.GetUserByEmailAddressQuery{
		EmailAddress: userEmailAddress,
	})
	if err != nil {
		pgUow.RevertChanges()
		return
	}

	var userIdDto uuid.UUID
	if userDto != nil {
		userIdDto = userDto.Id
	} else {
		userIdDto, err = authAppService.Register(authappsrv.RegisterCommand{EmailAddress: userEmailAddress})
		if err != nil {
			pgUow.RevertChanges()
			return
		}

		// TODO - Remove this two phase commits across contexts by using integration events
		if err := worldAccountAppService.CreateWorldAccount(worldaccountappsrv.CreateWorldAccountCommand{
			UserId: userIdDto,
		}); err != nil {
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
			"%s/auth/sign-in-success/?access_token=%v&client_redirect_path=%v",
			clientUrl,
			accessToken,
			state.ClientRedirectPath,
		),
	)
}
