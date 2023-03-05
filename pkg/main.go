package main

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/intevent/redisinteventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/cassandra"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/socket/gamesocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	intEventPublisher := redisinteventpublisher.New()

	itemRepo, err := postgres.NewItemRepo()
	if err != nil {
		panic(err)
	}
	playerRepo := memrepo.NewPlayerMemRepo()
	worldRepo, err := postgres.NewWorldRepo()
	if err != nil {
		panic(err)
	}
	unitRepo, err := cassandra.NewUnitRepo()
	if err != nil {
		panic(err)
	}
	userRepo, err := postgres.NewUserRepo()
	if err != nil {
		panic(err)
	}
	userId, _ := usermodel.ParseUserIdVo("d169faa5-c078-42c2-8a42-cd1d43558c7b")
	newUser := usermodel.NewUnitAgg(userId, "dumdumgenius@gmail.com", "DumDumGenius")
	err = userRepo.Add(newUser)
	if err != nil {
		// panic(err)
	}

	gameSocketAppService := gamesocketappservice.NewService(intEventPublisher, worldRepo, playerRepo, unitRepo, itemRepo)
	gameSocketApiController := gamesocket.NewController(gameSocketAppService)

	err = gameSocketAppService.CreateWorld(userId.String())
	if err != nil {
		fmt.Println(err)
	}

	router.Static("/assets", "./pkg/interface/transport/assets")
	router.Group("/ws/game").GET("/", gameSocketApiController.HandleGameConnection)
	router.Run()

}
