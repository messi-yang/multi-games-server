package authhttphandler

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	game_pgrepository "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleauthinfraservice"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	googleAuthInfraService := googleauthinfraservice.NewService()

	userRepository, err := pgrepo.NewUserRepository()
	if err != nil {
		panic(err)
	}
	identityService := identitydomainsrv.NewService(userRepository)
	identityAppService := identityappsrv.NewService(userRepository, identityService, os.Getenv("AUTH_SECRET"))

	gamerRepository, err := game_pgrepository.NewGamerRepository()
	if err != nil {
		panic(err)
	}
	gamerAppService := gamerappsrv.NewService(gamerRepository)

	httpHandler := newHttpHandler(googleAuthInfraService, identityAppService, gamerAppService)

	routerGroup := router.Group("/api/auth")
	routerGroup.GET("/oauth2/google", httpHandler.goToGoogleAuthUrl)
	routerGroup.GET("/oauth2/google/redirect", httpHandler.googleAuthCallback)
}
