package gamesockethandler

import (
	"net/http"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Setup(router *gin.Engine) {
	itemRepo, err := pgrepo.NewItemRepo()
	if err != nil {
		panic(err)
	}
	playerRepo := memrepo.NewPlayerMemRepo()
	worldRepo, err := pgrepo.NewWorldRepo()
	if err != nil {
		panic(err)
	}
	unitRepo, err := pgrepo.NewUnitRepo()
	if err != nil {
		panic(err)
	}
	gameDomainService := gamedomainsrv.NewService(worldRepo, playerRepo, unitRepo, itemRepo)
	gameAppService := gameappsrv.NewService(
		worldRepo, playerRepo, unitRepo, itemRepo, gameDomainService,
	)
	httpHandler := newHttpHandler(gameAppService, websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})

	routerGroup := router.Group("/ws/game")
	routerGroup.GET("/", httpHandler.gameConnection)
}
