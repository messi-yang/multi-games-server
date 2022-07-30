package computeallgamesusecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
)

type useCase struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
	gameComputeEvent   gamecomputedevent.GameComputedEvent
}

func New(gameRoomRepository gameroomrepository.GameRoomRepository, gameComputeEvent gamecomputedevent.GameComputedEvent) *useCase {
	return &useCase{
		gameRoomRepository: gameRoomRepository,
		gameComputeEvent:   gameComputeEvent,
	}
}

func (uc *useCase) Execute() {
	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	gameRooms := uc.gameRoomRepository.GetAll()
	for _, gameRoom := range gameRooms {
		gameRoomService.GenerateNextUnitMap(gameRoom.GetGameId())
		uc.gameComputeEvent.Publish(gameRoom.GetGameId())
	}
}
