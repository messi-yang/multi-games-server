package computeallgamesusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
)

type useCase struct {
	gameRoomService  gameroomservice.GameRoomService
	gameComputeEvent gamecomputedevent.GameComputedEvent
}

func NewUseCase(gameRoomService gameroomservice.GameRoomService, gameComputeEvent gamecomputedevent.GameComputedEvent) *useCase {
	return &useCase{
		gameRoomService:  gameRoomService,
		gameComputeEvent: gameComputeEvent,
	}
}

func (uc *useCase) Execute() {
	gameRooms := uc.gameRoomService.GetAllGameRooms()
	for _, gameRoom := range gameRooms {
		uc.gameRoomService.GenerateNextGameUnitMatrix(gameRoom.GetGameId())
		uc.gameComputeEvent.Publish(gameRoom.GetGameId())
	}
}
