package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/http/gamehandler"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	gameApplicationService, _ := applicationservice.NewGameApplicationService(applicationservice.WithGameService())

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		GameApplicationService: gameApplicationService,
		NotificationPublisher:  commonredis.NewRedisNotificationPublisher(),
	}))

	router.Run()
}
