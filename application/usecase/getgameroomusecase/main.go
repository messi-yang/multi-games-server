package getgameroomusecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/gameroomdto"
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

func (uc *useCase) Execute(gameId uuid.UUID) (gameroomdto.GameRoomDTO, error) {
	gameRoom, err := uc.gameRoomRepository.Get(gameId)
	if err != nil {
		return gameroomdto.GameRoomDTO{}, err
	}
	return gameroomdto.ToDTO(gameRoom), nil
}
