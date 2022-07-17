package reviveunitsusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/coordinatesupdatedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/google/uuid"
)

type useCase struct {
	gameRoomRepository      gameroomrepository.GameRoomRepository
	coordinatesUpdatedEvent coordinatesupdatedevent.CoordinatesUpdatedEvent
}

func New(gameRoomRepository gameroomrepository.GameRoomRepository, coordinatesUpdatedEvent coordinatesupdatedevent.CoordinatesUpdatedEvent) *useCase {
	return &useCase{
		gameRoomRepository:      gameRoomRepository,
		coordinatesUpdatedEvent: coordinatesUpdatedEvent,
	}
}

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) error {
	coordinates := coordinatedto.FromDTOList(coordinateDTOs)

	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	gameRoomService.ReviveUnits(gameId, coordinates)

	uc.coordinatesUpdatedEvent.Publish(gameId, coordinateDTOs)

	return nil
}
