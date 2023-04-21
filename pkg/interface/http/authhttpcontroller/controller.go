package authhttpcontroller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleoauthinfraservice"
	"github.com/gin-gonic/gin"
)

func goToGoogleAuthUrlHandler(c *gin.Context) {
	googleOauthInfraService := googleoauthinfraservice.NewService()
	authUrl := googleOauthInfraService.GenerateAuthUrl(googleoauthinfraservice.GenerateAuthUrlCommand{})
	c.Redirect(http.StatusFound, authUrl)
}

func googleAuthCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	googleOauthInfraService := googleoauthinfraservice.NewService()
	userEmailAddress, err := googleOauthInfraService.GetUserEmailAddress(googleoauthinfraservice.GetUserEmailAddressQuery{
		Code: code,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	clientUrl := os.Getenv("CLIENT_URL")
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/?email_address=%s", clientUrl, userEmailAddress))
}
