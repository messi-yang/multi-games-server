package httprouter

import (
	"net/http"
	"strings"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/authhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/gamerhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/gamesockethandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/itemhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/userhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/worldhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Authorization", "Origin"}
	router.Use(cors.New(corsConfig))

	router.Static("/asset", "./asset")

	authorizeTokenMiddleware := func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			ctx.String(http.StatusUnauthorized, "Token is not found in Authorization header")
			return
		}
		authToken := strings.Split(authorizationHeader, " ")[1]

		pgUow := pguow.NewUow()

		authAppService := providedependency.ProvideAuthAppService(pgUow)
		userId, err := authAppService.Validate(authToken)
		if err != nil {
			pgUow.RevertChanges()
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		pgUow.SaveChanges()
		httputil.SetUserId(ctx, userId)
		ctx.Next()
	}

	authHttpHandler := authhttphandler.NewHttpHandler()
	authRouterGroup := router.Group("/api/auth")
	authRouterGroup.GET("/oauth2/google", authHttpHandler.GoToGoogleAuthUrl)
	authRouterGroup.GET("/oauth2/google/redirect", authHttpHandler.HandleGoogleAuthCallback)

	userHttpHandler := userhttphandler.NewHttpHandler()
	userRouterGroup := router.Group("/api/user")
	userRouterGroup.GET("/me", authorizeTokenMiddleware, userHttpHandler.GetMyUser)

	gamerHttpHandler := gamerhttphandler.NewHttpHandler()
	gamersRouterGroup := router.Group("/api/gamers")
	gamersRouterGroup.GET("/", gamerHttpHandler.QueryGamers)

	worldHttphandler := worldhttphandler.NewHttpHandler()
	worldRouterGroup := router.Group("/api/worlds")
	worldRouterGroup.GET("/:worldId", worldHttphandler.GetWorld)
	worldRouterGroup.GET("/", worldHttphandler.QueryWorlds)
	worldRouterGroup.POST("/", authorizeTokenMiddleware, worldHttphandler.CreateWorld)
	worldRouterGroup.PATCH("/:worldId", authorizeTokenMiddleware, worldHttphandler.UpdateWorld)

	itemHttpHandler := itemhttphandler.NewHttpHandler()
	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", itemHttpHandler.QueryItems)

	redisServerMessageMediator := redisservermessagemediator.NewMediator()
	gameSocketHandler := gamesockethandler.NewHttpHandler(redisServerMessageMediator)
	gameRouterGroup := router.Group("/ws/game")
	gameRouterGroup.GET("/", gameSocketHandler.GameConnection)

	return router.Run()
}
