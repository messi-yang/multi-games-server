package creategameusecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
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

func (uc *useCase) Execute() (gameId uuid.UUID, err error) {
	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	size := config.GetConfig().GetGameMapSize()
	mapSize := valueobject.NewMapSize(size, size)
	gameRoom, err := gameRoomService.CreateGameRoom(mapSize)
	if err != nil {
		return uuid.UUID{}, err
	}
	return gameRoom.GetGameId(), nil
}
