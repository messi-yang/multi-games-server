package authhttpcontroller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/service/identityappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleauthinfraservice"
	"github.com/gin-gonic/gin"
)

func goToGoogleAuthUrlHandler(c *gin.Context) {
	googleAuthInfraService := provideGoogleOauthInfraService()
	authUrl := googleAuthInfraService.GenerateAuthUrl(googleauthinfraservice.GenerateAuthUrlCommand{})
	c.Redirect(http.StatusFound, authUrl)
}

func googleAuthCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	googleAuthInfraService := provideGoogleOauthInfraService()
	identityAppService, err := provideIdentityAppService()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userEmailAddress, err := googleAuthInfraService.GetUserEmailAddress(googleauthinfraservice.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user, err := identityAppService.LoginOrRegister(
		identityappservice.LoginOrRegisterCommand{EmailAddress: userEmailAddress},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(user)
	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/?access_token=%v", clientUrl, user))
}
