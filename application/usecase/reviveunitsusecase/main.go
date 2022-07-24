package reviveunitsusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/patterndto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/coordinatesupdatedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
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

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTO coordinatedto.CoordinateDTO, patternDTO patterndto.PatternDTO) error {
	coordinates := make([]valueobject.Coordinate, 0)
	for relativeX, rowsInPattern := range patternDTO {
		for relativeY, isTruthy := range rowsInPattern {
			if isTruthy {
				coordinates = append(
					coordinates,
					valueobject.NewCoordinate(
						coordinateDTO.X+relativeX,
						coordinateDTO.Y+relativeY,
					),
				)
			}
		}
	}

	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	gameRoomService.ReviveUnits(gameId, coordinates)

	coordinateDTOs := coordinatedto.ToDTOList(coordinates)
	uc.coordinatesUpdatedEvent.Publish(gameId, coordinateDTOs)

	return nil
}
