package mainserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/notification/commonredisnotification"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/interface/http/gamehandler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	gameApplicationService, _ := applicationservice.NewGameApplicationService(applicationservice.WithGameService())

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		GameApplicationService: gameApplicationService,
		NotificationPublisher:  commonredisnotification.NewRedisNotificationPublisher(),
	}))

	router.Run()
}
