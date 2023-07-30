package httprouter

import (
	"strings"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/authhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/itemhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/userhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/worldaccounthttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/worldhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/worldjourneyhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/worldmemberhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:    []string{"Authorization", "Origin", "Content-Type"},
	}))

	router.Static("/asset", "./asset")

	parseAccessTokenMiddleware := func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			ctx.Next()
			return
		}
		authToken := strings.Split(authorizationHeader, " ")[1]

		pgUow := pguow.NewDummyUow()

		authAppService := providedependency.ProvideAuthAppService(pgUow)
		userId, err := authAppService.Validate(authToken)
		if err != nil {
			ctx.Next()
			return
		}

		httpsession.SetAuthrorizedUserId(ctx, userId)
		ctx.Next()
	}

	authHttpHandler := authhttphandler.NewHttpHandler()
	authRouterGroup := router.Group("/api/auth")
	authRouterGroup.GET("/oauth2/google", authHttpHandler.GoToGoogleAuthUrl)
	authRouterGroup.GET("/oauth2/google/redirect", authHttpHandler.HandleGoogleAuthCallback)

	userHttpHandler := userhttphandler.NewHttpHandler()
	userRouterGroup := router.Group("/api/users")
	userRouterGroup.GET("/me", parseAccessTokenMiddleware, userHttpHandler.GetMyUser)
	userRouterGroup.PATCH("/me", parseAccessTokenMiddleware, userHttpHandler.UpdateMyUser)

	worldAccountHttpHandler := worldaccounthttphandler.NewHttpHandler()
	worldAccountsRouterGroup := router.Group("/api/world-accounts")
	worldAccountsRouterGroup.GET("/", worldAccountHttpHandler.QueryWorldAccounts)

	worldHttpHandler := worldhttphandler.NewHttpHandler()
	worldRouterGroup := router.Group("/api/worlds")
	worldRouterGroup.GET("/:worldId", worldHttpHandler.GetWorld)
	worldRouterGroup.GET("/", worldHttpHandler.QueryWorlds)
	worldRouterGroup.GET("/mine", parseAccessTokenMiddleware, worldHttpHandler.GetMyWorlds)
	worldRouterGroup.POST("/", parseAccessTokenMiddleware, worldHttpHandler.CreateWorld)
	worldRouterGroup.PATCH("/:worldId", parseAccessTokenMiddleware, worldHttpHandler.UpdateWorld)
	worldRouterGroup.DELETE("/:worldId", parseAccessTokenMiddleware, worldHttpHandler.DeleteWorld)

	worldMemberHttpHandler := worldmemberhttphandler.NewHttpHandler()
	worldRouterGroup.GET("/:worldId/members", parseAccessTokenMiddleware, worldMemberHttpHandler.GetWorldMembers)

	itemHttpHandler := itemhttphandler.NewHttpHandler()
	itemRouterGroup := router.Group("/api/items")
	itemRouterGroup.GET("/", itemHttpHandler.QueryItems)

	redisServerMessageMediator := redisservermessagemediator.NewMediator()
	worldJourneyHandler := worldjourneyhandler.NewHttpHandler(redisServerMessageMediator)
	worldJourneyGroup := router.Group("/api/world-journey")
	worldJourneyGroup.GET("/", worldJourneyHandler.StartJourney)

	return router.Run()
}
