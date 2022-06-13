package computeallgamesusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gamecomputeevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
)

type useCase struct {
	gameRoomService  gameroomservice.GameRoomService
	gameComputeEvent gamecomputeevent.GameComputeEvent
}

func NewUseCase(gameRoomService gameroomservice.GameRoomService, gameComputeEvent gamecomputeevent.GameComputeEvent) *useCase {
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
