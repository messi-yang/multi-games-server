package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/interface/http/gameroomhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
	})

	router.Group("/ws/game").GET("/", gameroomhandler.NewHandler(gameroomhandler.HandlerConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
		IntegrationEventBus:        redisIntegrationEventBus,
	}))

	router.Run()
}
