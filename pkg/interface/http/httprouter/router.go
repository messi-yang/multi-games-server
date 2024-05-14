package httprouter

import (
	"strings"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/authhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/embedunithttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/itemhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/linkunithttphandler"
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

	parseHttpAccessTokenMiddleware := func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			ctx.Next()
			return
		}
		accessToken := strings.Split(authorizationHeader, " ")[1]

		validateAccessTokenUseCase := usecase.ProvideValidateAccessTokenUseCase()
		userIdDto, err := validateAccessTokenUseCase.Execute(accessToken)
		if err != nil {
			ctx.Next()
			return
		}

		httpsession.SetAuthrorizedUserId(ctx, userIdDto)
		ctx.Next()
	}

	parseSocketAccessTokenMiddleware := func(ctx *gin.Context) {
		accessToken := ctx.Request.URL.Query().Get("access-token")

		validateAccessTokenUseCase := usecase.ProvideValidateAccessTokenUseCase()
		userIdDto, err := validateAccessTokenUseCase.Execute(accessToken)
		if err != nil {
			ctx.Next()
			return
		}

		httpsession.SetAuthrorizedUserId(ctx, userIdDto)
		ctx.Next()
	}

	authHttpHandler := authhttphandler.NewHttpHandler()
	authRouterGroup := router.Group("/api/auth")
	authRouterGroup.GET("/oauth2/google", authHttpHandler.RedirectToGoogleOauthUrl)
	authRouterGroup.GET("/oauth2/google/redirect", authHttpHandler.HandleGoogleOauthCallback)

	userHttpHandler := userhttphandler.NewHttpHandler()
	userRouterGroup := router.Group("/api/users")
	userRouterGroup.Use(parseHttpAccessTokenMiddleware)
	userRouterGroup.GET("/me", userHttpHandler.GetMyUser)
	userRouterGroup.PATCH("/me", userHttpHandler.UpdateMyUser)

	worldAccountHttpHandler := worldaccounthttphandler.NewHttpHandler()
	worldAccountsRouterGroup := router.Group("/api/world-accounts")
	worldAccountsRouterGroup.Use(parseHttpAccessTokenMiddleware)
	worldAccountsRouterGroup.GET("/", worldAccountHttpHandler.QueryWorldAccounts)

	worldHttpHandler := worldhttphandler.NewHttpHandler()
	worldRouterGroup := router.Group("/api/worlds")
	worldRouterGroup.Use(parseHttpAccessTokenMiddleware)

	worldRouterGroup.GET("/:worldId", worldHttpHandler.GetWorld)
	worldRouterGroup.GET("/", worldHttpHandler.QueryWorlds)
	worldRouterGroup.GET("/mine", worldHttpHandler.GetMyWorlds)
	worldRouterGroup.POST("/", worldHttpHandler.CreateWorld)
	worldRouterGroup.PATCH("/:worldId", worldHttpHandler.UpdateWorld)
	worldRouterGroup.DELETE("/:worldId", worldHttpHandler.DeleteWorld)

	worldMemberHttpHandler := worldmemberhttphandler.NewHttpHandler()
	worldRouterGroup.GET("/:worldId/members", worldMemberHttpHandler.GetWorldMembers)

	itemHttpHandler := itemhttphandler.NewHttpHandler()
	itemRouterGroup := router.Group("/api/items")
	itemRouterGroup.Use(parseHttpAccessTokenMiddleware)
	itemRouterGroup.GET("/", itemHttpHandler.QueryItems)
	itemRouterGroup.GET("/with-ids", itemHttpHandler.GetItemsWithIds)

	linkUnitHttpHandler := linkunithttphandler.NewHttpHandler()
	linkUnitRouterGroup := router.Group("/api/link-units")
	linkUnitRouterGroup.GET("/:id", linkUnitHttpHandler.GetLinkUnitUrl)

	embedUnitHttpHandler := embedunithttphandler.NewHttpHandler()
	embedUnitRouterGroup := router.Group("/api/embed-units")
	embedUnitRouterGroup.GET("/:id", embedUnitHttpHandler.GetEmbedUnitEmbedCode)

	redisServerMessageMediator := redisservermessagemediator.NewMediator()
	worldJourneyHandler := worldjourneyhandler.NewHttpHandler(redisServerMessageMediator)
	worldJourneyGroup := router.Group("/api/world-journey")
	worldJourneyGroup.Use(parseSocketAccessTokenMiddleware)
	worldJourneyGroup.GET("/", worldJourneyHandler.StartJourney)

	return router.Run()
}
