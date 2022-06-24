package reviveunitsusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/unitsupdatedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/google/uuid"
)

type useCase struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
	unitsUpdatedEvent  unitsupdatedevent.UnitsUpdatedEvent
}

func New(gameRoomRepository gameroomrepository.GameRoomRepository, unitsUpdatedEvent unitsupdatedevent.UnitsUpdatedEvent) *useCase {
	return &useCase{
		gameRoomRepository: gameRoomRepository,
		unitsUpdatedEvent:  unitsUpdatedEvent,
	}
}

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) error {
	coordinates := coordinatedto.FromDTOList(coordinateDTOs)

	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	gameRoomService.ReviveUnits(gameId, coordinates)

	uc.unitsUpdatedEvent.Publish(gameId, coordinateDTOs)

	return nil
}
