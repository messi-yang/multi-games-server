package httprouter

import (
	"net/http"
	"strings"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/authhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/gamerhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/gamesockethandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/itemhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httphandler/worldhttphandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Run() error {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.Static("/asset", "./asset")

	googleAuthInfraService := provideGoogleAuthInfraService()
	identityAppService, err := provideIdentityAppService()
	if err != nil {
		return err
	}
	gamerAppService, err := provideGamerAppService()
	if err != nil {
		return err
	}
	worldAppService, err := provideWorldAppService()
	if err != nil {
		return err
	}
	itemAppService, err := provideItemAppService()
	if err != nil {
		return err
	}
	gameAppService, err := provideGameAppService()
	if err != nil {
		return err
	}

	authorizeToken := func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			ctx.String(http.StatusUnauthorized, "Token is not found in Authorization header")
			return
		}
		authToken := strings.Split(authorizationHeader, " ")[1]
		userId, err := identityAppService.Validate(authToken)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
		}

		httputil.SetUserId(ctx, userId)

		ctx.Next()
	}

	authHttpHandler := authhttphandler.NewHttpHandler(googleAuthInfraService, identityAppService, gamerAppService)
	authRouterGroup := router.Group("/api/auth")
	authRouterGroup.GET("/oauth2/google", authHttpHandler.GoToGoogleAuthUrl)
	authRouterGroup.GET("/oauth2/google/redirect", authHttpHandler.HandleGoogleAuthCallback)

	gamerHttpHandler := gamerhttphandler.NewHttpHandler(gamerAppService)
	gamersRouterGroup := router.Group("/api/gamers")
	gamersRouterGroup.GET("/", gamerHttpHandler.QueryGamers)

	worldHttphandler := worldhttphandler.NewHttpHandler(identityAppService, worldAppService, gamerAppService)
	worldRouterGroup := router.Group("/api/worlds")
	worldRouterGroup.GET("/:worldId", worldHttphandler.GetWorld)
	worldRouterGroup.GET("/", worldHttphandler.QueryWorlds)
	worldRouterGroup.POST("/", authorizeToken, worldHttphandler.CreateWorld)

	itemHttpHandler := itemhttphandler.NewHttpHandler(itemAppService)
	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", itemHttpHandler.QueryItems)

	websocketUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	gameSocketHandler := gamesockethandler.NewHttpHandler(gameAppService, websocketUpgrader)
	gameRouterGroup := router.Group("/ws/game")
	gameRouterGroup.GET("/", gameSocketHandler.GameConnection)

	return router.Run()
}
