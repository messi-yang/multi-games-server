package commonserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/commonserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/commonserver/interface/http/gamehandler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	liveGameApplicationService, _ := applicationservice.NewLiveGameApplicationService(applicationservice.WithLiveGameService())

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		LiveGameApplicationService: liveGameApplicationService,
	}))

	router.Run()
}
