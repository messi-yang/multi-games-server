package authhttphandler

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappservice"
	game_pg_repository "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/service/identityappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/persistence/pgrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/infrastructure/service/googleauthinfraservice"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	googleAuthInfraService := googleauthinfraservice.NewService()

	userRepository, err := pgrepository.NewUserRepository()
	if err != nil {
		panic(err)
	}
	identityService := service.NewIdentityService(userRepository)
	identityAppService := identityappservice.NewService(userRepository, identityService, os.Getenv("AUTH_SECRET"))

	gamerRepository, err := game_pg_repository.NewGamerRepository()
	if err != nil {
		panic(err)
	}
	gamerAppService := gamerappservice.NewService(gamerRepository)

	httpHandler := newHttpHandler(googleAuthInfraService, identityAppService, gamerAppService)

	routerGroup := router.Group("/api/auth")
	routerGroup.GET("/oauth2/google", httpHandler.goToGoogleAuthUrl)
	routerGroup.GET("/oauth2/google/redirect", httpHandler.googleAuthCallback)
}
