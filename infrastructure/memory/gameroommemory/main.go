package gameroommemory

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameRoomMemoryRepository interface {
	Create(aggregate.GameRoom) error
	UpdateGameUnit(uuid.UUID, valueobject.Coordinate, valueobject.GameUnit) error
	UpdateGameUnitMatrix(uuid.UUID, [][]valueobject.GameUnit) error
	Get(uuid.UUID) (aggregate.GameRoom, error)
	GetAll() []aggregate.GameRoom
}

type gameRoomMemoryRepositoryImpl struct {
	gameRoomMap map[uuid.UUID]aggregate.GameRoom
}

var gameRoomMemoryRepository GameRoomMemoryRepository

func NewGameRoomMemoryRepository() GameRoomMemoryRepository {
	if gameRoomMemoryRepository == nil {
		gameRoomMemoryRepository = &gameRoomMemoryRepositoryImpl{
			gameRoomMap: make(map[uuid.UUID]aggregate.GameRoom),
		}
		return gameRoomMemoryRepository
	} else {
		return gameRoomMemoryRepository
	}
}

func (gmi *gameRoomMemoryRepositoryImpl) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	gameRoom, exists := gmi.gameRoomMap[id]
	if !exists {
		return aggregate.GameRoom{}, gameroomrepository.ErrGameRoomNotFound
	}
	return gameRoom, nil
}

func (gmi *gameRoomMemoryRepositoryImpl) GetAll() []aggregate.GameRoom {
	gameRooms := make([]aggregate.GameRoom, 0)
	for _, gameRoom := range gmi.gameRoomMap {
		gameRooms = append(gameRooms, gameRoom)
	}
	return gameRooms
}

func (gmi *gameRoomMemoryRepositoryImpl) UpdateGameUnit(gameId uuid.UUID, coordinate valueobject.Coordinate, gameUnit valueobject.GameUnit) error {
	gameRoom := gmi.gameRoomMap[gameId]
	gameRoom.UpdateGameUnit(coordinate, gameUnit)

	return nil
}

func (gmi *gameRoomMemoryRepositoryImpl) UpdateGameUnitMatrix(gameId uuid.UUID, gameUnitMatrix [][]valueobject.GameUnit) error {
	gameRoom := gmi.gameRoomMap[gameId]
	gameRoom.UpdateGameUnitMatrix(gameUnitMatrix)

	return nil
}

func (gmi *gameRoomMemoryRepositoryImpl) Create(gameRoom aggregate.GameRoom) error {
	gmi.gameRoomMap[gameRoom.GetGameId()] = gameRoom

	return nil
}
