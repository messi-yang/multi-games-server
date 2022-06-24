package computeallgamesusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
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
		gameRoomService.GenerateNextUnitMatrix(gameRoom.GetGameId())
		uc.gameComputeEvent.Publish(gameRoom.GetGameId())
	}
}
