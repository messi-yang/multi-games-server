package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/repositorymemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/jobs/tickunitmapjob"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

func Start() {
	gameRoomRepositoryMemory := repositorymemory.NewGameRoomRepositoryMemory()
	size := config.GetConfig().GetGameMapSize()
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{GameRoomRepository: gameRoomRepositoryMemory},
	)
	newGameRoomId, err := gameRoomApplicationService.CreateRoom(size, size)
	if err != nil {
		panic(err.Error())
	}
	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	redisInfrastructureService.Set("game_id", []byte(newGameRoomId.String()))

	gameRoomJob := tickunitmapjob.GetJob()
	gameRoomJob.Start()

	integrationeventhandler.HandleGameRoomIntegrationEvent()
}
