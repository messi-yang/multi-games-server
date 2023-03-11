package main

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/client/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/redisinteventpublisher"
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

	redisClient := redisclient.New()
	intEventPublisher := redisinteventpublisher.New(redisClient)

	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		panic(err)
	}
	playerRepository := memrepo.NewPlayerMemRepository()
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		panic(err)
	}
	unitRepository, err := cassandra.NewUnitRepository()
	if err != nil {
		panic(err)
	}
	userRepository, err := postgres.NewUserRepository()
	if err != nil {
		panic(err)
	}
	userId, _ := usermodel.ParseUserIdVo("d169faa5-c078-42c2-8a42-cd1d43558c7b")
	newUser := usermodel.NewUnitAgg(userId, "dumdumgenius@gmail.com", "DumDumGenius")
	err = userRepository.Add(newUser)
	if err != nil {
		// panic(err)
	}

	gameSocketAppService := gamesocketappservice.NewService(intEventPublisher, worldRepository, playerRepository, unitRepository, itemRepository)
	gameSocketApiController := gamesocket.NewController(gameSocketAppService, redisClient)

	err = gameSocketAppService.CreateWorld(userId.String())
	if err != nil {
		fmt.Println(err)
	}

	router.Static("/asset", "./pkg/interface/transport/asset")
	router.Group("/ws/game").GET("/", gameSocketApiController.HandleGameConnection)
	router.Run()

}
