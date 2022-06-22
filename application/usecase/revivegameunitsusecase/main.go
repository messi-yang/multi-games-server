package revivegameunitsusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gameunitsupdatedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/google/uuid"
)

type useCase struct {
	gameRoomRepository    gameroomrepository.GameRoomRepository
	gameUnitsUpdatedEvent gameunitsupdatedevent.GameUnitsUpdatedEvent
}

func NewUseCase(gameRoomRepository gameroomrepository.GameRoomRepository, gameUnitsUpdatedEvent gameunitsupdatedevent.GameUnitsUpdatedEvent) *useCase {
	return &useCase{
		gameRoomRepository:    gameRoomRepository,
		gameUnitsUpdatedEvent: gameUnitsUpdatedEvent,
	}
}

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) error {
	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	coordinates := coordinatedto.FromDTOs(coordinateDTOs)

	for _, coord := range coordinates {
		gameRoomService.ReviveGameUnit(gameId, coord)
	}

	uc.gameUnitsUpdatedEvent.Publish(gameId, coordinates)

	return nil
}
