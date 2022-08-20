package creategameroomusecase

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
)

type useCase struct {
	gameRoomRepository gameroomrepository.GameRoomRepository
}

func New(gameRoomRepository gameroomrepository.GameRoomRepository) *useCase {
	return &useCase{
		gameRoomRepository: gameRoomRepository,
	}
}

func (uc *useCase) Execute() (aggregate.GameRoom, error) {
	gameRoomService := gameroomservice.NewGameRoomService(uc.gameRoomRepository)
	size := config.GetConfig().GetGameMapSize()
	mapSize := valueobject.NewMapSize(size, size)
	gameRoom, err := gameRoomService.CreateGameRoom(mapSize)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return *gameRoom, nil
}
