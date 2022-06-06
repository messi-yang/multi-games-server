package updateallgamesusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gameupdateevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
)

type useCase struct {
	gameRoomService gameroomservice.GameRoomService
	gameUpdateEvent *gameupdateevent.GameUpdateEvent
}

func NewUseCase(gameRoomService gameroomservice.GameRoomService, gameUpdateEvent *gameupdateevent.GameUpdateEvent) *useCase {
	return &useCase{
		gameRoomService: gameRoomService,
		gameUpdateEvent: gameUpdateEvent,
	}
}

func (uc *useCase) Execute() {
	gameRooms := uc.gameRoomService.GetAllGameRooms()
	for _, gameRoom := range gameRooms {
		uc.gameRoomService.GenerateNextGameUnitMatrix(gameRoom.GetGameId())
		uc.gameUpdateEvent.Publish(gameRoom.GetGameId())
	}
}
