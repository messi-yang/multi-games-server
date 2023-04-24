package authhttphandler

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	game_pgrepository "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/service/identityappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/service/googleauthinfrasrv"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	googleAuthInfraService := googleauthinfrasrv.NewService()

	userRepo, err := pgrepo.NewUserRepo()
	if err != nil {
		panic(err)
	}
	identityService := identitydomainsrv.NewService(userRepo)
	identityAppService := identityappsrv.NewService(userRepo, identityService, os.Getenv("AUTH_SECRET"))

	gamerRepo, err := game_pgrepository.NewGamerRepo()
	if err != nil {
		panic(err)
	}
	gamerAppService := gamerappsrv.NewService(gamerRepo)

	httpHandler := newHttpHandler(googleAuthInfraService, identityAppService, gamerAppService)

	routerGroup := router.Group("/api/auth")
	routerGroup.GET("/oauth2/google", httpHandler.goToGoogleAuthUrl)
	routerGroup.GET("/oauth2/google/redirect", httpHandler.googleAuthCallback)
}
