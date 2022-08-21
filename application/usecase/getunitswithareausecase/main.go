package getunitswithareausecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/google/uuid"
)

type useCase struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
}

func New(gameRoomRepository gameroomrepository.GameRoomRepository) *useCase {
	return &useCase{
		gameRoomRepository: gameRoomRepository,
	}
}

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, areaDTO areadto.AreaDTO) ([]coordinatedto.CoordinateDTO, []unitdto.UnitDTO, error) {
	gameRoom, err := uc.gameRoomRepository.Get(gameId)
	if err != nil {
		return nil, nil, err
	}

	coordinates := coordinatedto.FromDTOList(coordinateDTOs)
	area, err := areadto.FromDTO(areaDTO)
	if err != nil {
		return nil, nil, err
	}

	coordinatesInArea, err := gameRoom.FilterCoordinatesWithArea(coordinates, area)
	if err != nil {
		return nil, nil, err
	}

	units, err := gameRoom.GetUnitsWithCoordinates(coordinatesInArea)
	if err != nil {
		return nil, nil, err
	}

	return coordinatedto.ToDTOList(coordinatesInArea), unitdto.ToDTOList(units), nil
}
