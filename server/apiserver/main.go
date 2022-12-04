package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/http/gamehandler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	liveGameApplicationService, _ := service.NewLiveGameApplicationService(
		service.WithRedisNotificationPublisher(),
	)
	gameApplicationService, _ := service.NewGameApplicationService(
		service.WithGameService(),
	)

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		LiveGameApplicationService: liveGameApplicationService,
		GameApplicationService:     gameApplicationService,
	}))

	router.Run()
}
