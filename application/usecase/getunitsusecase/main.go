package getunitsusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/unitdto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
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

func (uc *useCase) Execute(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) ([]unitdto.UnitDTO, error) {
	coordinates := coordinatedto.FromDTOList(coordinateDTOs)

	gameRoom, err := uc.gameRoomRepository.Get(gameId)
	if err != nil {
		return make([]unitdto.UnitDTO, 0), err
	}

	units, err := gameRoom.GetUnitsWithCoordinates(coordinates)
	if err != nil {
		return make([]unitdto.UnitDTO, 0), err
	}

	return unitdto.ToDTOList(units), nil
}
