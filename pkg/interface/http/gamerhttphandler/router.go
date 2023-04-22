package gamerhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepository"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	gamerRepository, err := pgrepository.NewGamerRepository()
	if err != nil {
		if err != nil {
			panic(err)
		}
	}
	gamerAppService := gamerappservice.NewService(gamerRepository)
	httpHandler := newHttpHandler(gamerAppService)

	routerGroup := router.Group("/api/gamers")
	routerGroup.GET("/", httpHandler.queryGamers)
}
