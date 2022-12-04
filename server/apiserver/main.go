package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/httpcontroller/livegamecontroller"
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

	router.Group("/ws/game").GET("/", livegamecontroller.NewController(livegamecontroller.Configuration{
		LiveGameApplicationService: liveGameApplicationService,
		GameApplicationService:     gameApplicationService,
	}))

	router.Run()
}
