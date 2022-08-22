package reviveunitsusecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/coordinatesupdatedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
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
	coordinates, err := coordinatedto.FromDTOList(coordinateDTOs)
	if err != nil {
		return err
	}

	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	err = gameRoomService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return err
	}

	uc.coordinatesUpdatedEvent.Publish(gameId, coordinateDTOs)

	return nil
}
