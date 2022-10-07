package gameclientcommunicator

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/interface/http/gameroomhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	integrationEventBusRedis := eventbusredis.NewIntegrationEventBusRedis(eventbusredis.IntegrationEventBusRedisCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
	})

	router.Group("/ws/game").GET("/", gameroomhandler.NewHandler(gameroomhandler.HandlerConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
		IntegrationEventBus:        integrationEventBusRedis,
	}))

	router.Run()
}
