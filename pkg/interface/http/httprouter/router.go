package httprouter

import (
	"strings"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/authhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/roomhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/roommemberhttphandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/roomservicehandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httphandler/userhttphandler"
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

	roomHttpHandler := roomhttphandler.NewHttpHandler()
	roomRouterGroup := router.Group("/api/rooms")
	roomRouterGroup.Use(parseHttpAccessTokenMiddleware)

	roomRouterGroup.GET("/:roomId", roomHttpHandler.GetRoom)
	roomRouterGroup.GET("/mine", roomHttpHandler.GetMyRooms)
	roomRouterGroup.POST("/", roomHttpHandler.CreateRoom)
	roomRouterGroup.PATCH("/:roomId", roomHttpHandler.UpdateRoom)
	roomRouterGroup.DELETE("/:roomId", roomHttpHandler.DeleteRoom)

	roomMemberHttpHandler := roommemberhttphandler.NewHttpHandler()
	roomRouterGroup.GET("/:roomId/members", roomMemberHttpHandler.GetRoomMembers)

	redisServerMessageMediator := redisservermessagemediator.NewMediator()
	roomServiceHandler := roomservicehandler.NewHttpHandler(redisServerMessageMediator)
	roomServiceGroup := router.Group("/api/room-service")
	roomServiceGroup.Use(parseSocketAccessTokenMiddleware)
	roomServiceGroup.GET("/", roomServiceHandler.StartService)

	return router.Run()
}
