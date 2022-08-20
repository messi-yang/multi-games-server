package getcoordinatesinareausecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
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

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, areaDTO areadto.AreaDTO) ([]coordinatedto.CoordinateDTO, error) {
	gameRoom, err := uc.gameRoomRepository.Get(gameId)
	if err != nil {
		return make([]coordinatedto.CoordinateDTO, 0), err
	}

	coordinates := coordinatedto.FromDTOList(coordinateDTOs)
	area := areadto.FromDTO(areaDTO)

	coordinatesInArea, err := gameRoom.FilterCoordinatesWithArea(coordinates, area)
	if err != nil {
		return nil, err
	}
	coordinateDTOsInArea := coordinatedto.ToDTOList(coordinatesInArea)

	return coordinateDTOsInArea, nil
}
