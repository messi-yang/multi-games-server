package getgameroomusecase

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/google/uuid"
)

type useCase struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
}

func NewUseCase(gameRoomRepository gameroomrepository.GameRoomRepository) *useCase {
	return &useCase{
		gameRoomRepository: gameRoomRepository,
	}
}

func (uc *useCase) Execute(gameId uuid.UUID) (aggregate.GameRoom, error) {
	gameRoom, err := uc.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return gameRoom, nil
}
